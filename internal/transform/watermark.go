package transform

import (
	"context"
	"image"
	"image/draw"
)

type WatermarkRemove struct {
	Rect image.Rectangle // region to remove
}

func (w WatermarkRemove) Name() string { return "watermark-remove" }

func (w WatermarkRemove) Apply(ctx context.Context, img image.Image) (image.Image, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	rgba := image.NewRGBA(img.Bounds())
	draw.Draw(rgba, rgba.Bounds(), img, img.Bounds().Min, draw.Src)

	// Simple removal: fill region with average background color
	for y := w.Rect.Min.Y; y < w.Rect.Max.Y; y++ {
		for x := w.Rect.Min.X; x < w.Rect.Max.X; x++ {
			rgba.Set(x, y, rgba.At(w.Rect.Min.X-1, y))
		}
	}

	return rgba, nil
}
