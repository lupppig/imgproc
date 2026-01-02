package cli

import (
	"errors"

	"github.com/lupppig/imgproc/internal/config"
)

func Validate(cfg *config.Config) error {
	if cfg.InputDir == "" {
		return errors.New("input directory is required")
	}
	if cfg.OutputDir == "" {
		return errors.New("output directory is required")
	}
	if cfg.Quality < 1 || cfg.Quality > 100 {
		return errors.New("quality must be between 1 and 100")
	}
	return nil
}
