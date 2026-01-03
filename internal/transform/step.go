package transform

import (
	"context"
	"image"
)

type Step interface {
	Name() string
	Apply(ctx context.Context, img image.Image) (image.Image, error)
}
