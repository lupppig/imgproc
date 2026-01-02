package app

import (
	"context"
)

type Flag interface {
	Run(ctx context.Context)
	Error() error
}
