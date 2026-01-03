package commands

import (
	"context"
	"os"
	"path/filepath"

	"github.com/lupppig/imgproc/internal/pipeline"
)

var supportedExt = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".webp": true,
}

func ProduceJobs(ctx context.Context, cfg ProcessConfig, jobs chan<- pipeline.ImageJob) error {
	info, err := os.Stat(cfg.InputDir)
	if err != nil {
		return err
	}

	if info.IsDir() {
		return walkDir(ctx, cfg, jobs)
	}

	return sendJob(ctx, cfg, cfg.InputDir, jobs)
}

func walkDir(ctx context.Context, cfg ProcessConfig, jobs chan<- pipeline.ImageJob) error {

	return filepath.WalkDir(cfg.InputDir, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if d.IsDir() {
			return nil
		}

		if !isSupported(path) {
			return nil
		}

		return sendJob(ctx, cfg, path, jobs)
	})
}

func sendJob(ctx context.Context, cfg ProcessConfig, input string, jobs chan<- pipeline.ImageJob) error {

	job := pipeline.ImageJob{
		Input:        input,
		Output:       outputPath(cfg.OutputDir, input),
		Format:       cfg.Format,
		ResizeWidth:  cfg.ResizeWidth,
		Quality:      cfg.Quality,
		AttemptsLeft: 3,
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case jobs <- job:
		return nil
	}
}
