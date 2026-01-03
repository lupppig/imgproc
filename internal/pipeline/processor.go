package pipeline

import (
	"context"
	"fmt"
	"image"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/lupppig/imgproc/internal/transform"
)

type ImageJob struct {
	Input        string
	Output       string
	Format       string
	ResizeWidth  int
	Quality      int
	AttemptsLeft int
	StripEXIF    bool
	Watermark    bool
}

var ThumbnailSizes = []int{1024, 256, 64}

func fanoutPath(base string, size int, format string) string {
	ext := "." + format
	name := strings.TrimSuffix(filepath.Base(base), filepath.Ext(base))
	return filepath.Join(
		filepath.Dir(base),
		fmt.Sprintf("%s_%d%s", name, size, ext),
	)
}

func ProcessImage(job ImageJob, progress func(processed int)) error {

	ctx := context.Background()

	img, inputFormat, err := transform.Decode(job.Input)
	if err != nil {
		return err
	}

	steps := []transform.Step{}

	if !job.StripEXIF {
		steps = append(steps, transform.OrientationFix{Path: job.Input})
	}

	if job.Watermark {
		steps = append(steps, transform.WatermarkRemove{
			Rect: image.Rect(10, 10, 200, 80),
		})
	}

	for _, step := range steps {
		img, err = step.Apply(ctx, img)
		if err != nil {
			return err
		}
	}

	format := job.Format
	if format == "" {
		format = inputFormat
	}

	sizes := []int{1024, 256, 64}
	sort.Sort(sort.Reverse(sort.IntSlice(sizes)))

	current := img

	for _, size := range sizes {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		resized, err := transform.Resize{Width: size}.Apply(ctx, current)
		if err != nil {
			return err
		}

		out := fanoutPath(job.Output, size, format)
		if err := os.MkdirAll(filepath.Dir(out), 0o755); err != nil {
			return err
		}

		if err := transform.Encode(out, resized, format, job.Quality); err != nil {
			return err
		}

		if progress != nil {
			progress(size)
		}

		current = resized
	}

	return nil
}
