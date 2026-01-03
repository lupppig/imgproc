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

func ProduceJobs(ctx context.Context, cfg ProcessConfig, jobs chan<- pipeline.ImageJob) (int, error) {
	info, err := os.Stat(cfg.InputDir)
	if err != nil {
		return 0, err
	}

	if info.IsDir() {
		return walkDir(ctx, cfg, jobs)
	}

	err = sendJob(ctx, cfg, cfg.InputDir, jobs)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func walkDir(ctx context.Context, cfg ProcessConfig, jobs chan<- pipeline.ImageJob) (int, error) {
	total := 0

	err := filepath.Walk(cfg.InputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		if err := sendJob(ctx, cfg, path, jobs); err != nil {
			return err
		}
		total++
		return nil
	})

	if err != nil {
		return total, err
	}

	return total, nil
}

func sendJob(ctx context.Context, cfg ProcessConfig, path string, jobs chan<- pipeline.ImageJob) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
	}

	job := pipeline.ImageJob{
		Input:        path,
		Output:       filepath.Join(cfg.OutputDir, filepath.Base(path)),
		Format:       cfg.Format,
		ResizeWidth:  cfg.ResizeWidth,
		Quality:      cfg.Quality,
		AttemptsLeft: 3, // retries
		StripEXIF:    cfg.StripEXIF,
		Watermark:    cfg.Watermark,
	}

	jobs <- job
	return nil
}
