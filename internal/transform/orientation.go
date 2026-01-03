package transform

import (
	"context"
	"image"
	"image/draw"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func readOrientation(path string) (int, error) {
	f, err := os.Open(path)
	if err != nil {
		return 1, err
	}
	defer f.Close()

	x, err := exif.Decode(f)
	if err != nil {
		return 1, nil // no EXIF = normal
	}

	tag, err := x.Get(exif.Orientation)
	if err != nil {
		return 1, nil
	}

	ori, err := tag.Int(0)
	if err != nil {
		return 1, nil
	}

	return ori, nil
}

func applyOrientation(img image.Image, orientation int) image.Image {
	b := img.Bounds()

	switch orientation {
	case 3: // 180
		dst := image.NewRGBA(b)
		draw.Draw(dst, b, img, b.Max, draw.Src)
		return dst

	case 6: // 90 CW
		dst := image.NewRGBA(image.Rect(0, 0, b.Dy(), b.Dx()))
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				dst.Set(b.Max.Y-y-1, x, img.At(x, y))
			}
		}
		return dst

	case 8: // 90 CCW
		dst := image.NewRGBA(image.Rect(0, 0, b.Dy(), b.Dx()))
		for y := b.Min.Y; y < b.Max.Y; y++ {
			for x := b.Min.X; x < b.Max.X; x++ {
				dst.Set(y, b.Max.X-x-1, img.At(x, y))
			}
		}
		return dst

	default:
		return img
	}
}

type OrientationFix struct {
	Path string
}

func (o OrientationFix) Name() string { return "orientation" }

func (o OrientationFix) Apply(ctx context.Context, img image.Image) (image.Image, error) {
	if o.Path == "" {
		return img, nil
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	ori, err := readOrientation(o.Path)
	if err != nil || ori == 1 {
		return img, nil
	}

	return applyOrientation(img, ori), nil
}
