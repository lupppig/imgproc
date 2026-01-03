package cli

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/lupppig/imgproc/internal/commands"
)

func Validate(cfg *commands.Config) error {
	if cfg.InputDir == "" {
		return errors.New("input directory is required")
	}
	if cfg.OutputDir == "" {
		return errors.New("output directory is required")
	}
	if cfg.Quality < 1 || cfg.Quality > 100 {
		return errors.New("quality must be between 1 and 100")
	}

	switch cfg.Format {
	case "jpeg", "png", "webp":
	default:
		return fmt.Errorf("--format must be one of: jpeg, png, webp")
	}

	if _, err := os.Stat(cfg.InputDir); err != nil {
		return fmt.Errorf("input directory provided does not exist: %s", cfg.InputDir)
	}

	if err := os.MkdirAll(filepath.Clean(cfg.OutputDir), 0755); err != nil {
		return fmt.Errorf("cannot create output directory: %w", err)
	}

	return nil
}
