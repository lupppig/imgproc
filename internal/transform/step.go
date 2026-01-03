package transform

import (
	"context"
	"image"

	"golang.org/x/image/draw"
)

type Step interface {
	Name() string
	Apply(ctx context.Context, img image.Image) (image.Image, error)
}

type Resize struct {
	Width int
}

func (r Resize) Name() string { return "resize" }

func (r Resize) Apply(ctx context.Context, img image.Image) (image.Image, error) {
	if r.Width <= 0 {
		return img, nil
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	b := img.Bounds()
	ratio := float64(r.Width) / float64(b.Dx())
	h := int(float64(b.Dy()) * ratio)

	dst := image.NewRGBA(image.Rect(0, 0, r.Width, h))
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, b, draw.Over, nil)

	return dst, nil
}



