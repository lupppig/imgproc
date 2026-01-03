package image_test

import (
	"image"
	"image/color"
	"image/jpeg"
	"os"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/lupppig/imgproc/internal/pipeline"
	"github.com/stretchr/testify/assert"
)

func createTempJPEG(t *testing.T, path string, w, h int) {
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	// Fill with red color
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			img.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		}
	}

	f, err := os.Create(path)
	assert.NoError(t, err)
	defer f.Close()

	err = jpeg.Encode(f, img, &jpeg.Options{Quality: 90})
	assert.NoError(t, err)
}

func TestProcessImage(t *testing.T) {

	tmpDir := t.TempDir()
	inputFile := filepath.Join(tmpDir, "input.jpg")
	outputFile := filepath.Join(tmpDir, "output.jpg")

	// Create a sample image
	createTempJPEG(t, inputFile, 200, 200)

	job := pipeline.ImageJob{
		Input:        inputFile,
		Output:       outputFile,
		Format:       "jpeg",
		ResizeWidth:  100,
		Quality:      80,
		AttemptsLeft: 3,
		StripEXIF:    true,
		Watermark:    true,
	}

	err := pipeline.ProcessImage(job, func(processed int) {

	})
	assert.NoError(t, err)

	for _, size := range pipeline.ThumbnailSizes {
		size_str := strconv.Itoa(size)
		outPath := filepath.Join(tmpDir,
			filepath.Base(job.Output[:len(job.Output)-len(filepath.Ext(job.Output))]+"_"+size_str+".jpeg"))
		_, err := os.Stat(outPath)
		assert.NoError(t, err, "thumbnail not created: %s", outPath)
	}
}
