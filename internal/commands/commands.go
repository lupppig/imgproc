package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"sync"
	"sync/atomic"

	"github.com/lupppig/imgproc/internal/pipeline"
)

type Command interface {
	Run(ctx context.Context) error
}

type HelpCommand struct{}

func (h *HelpCommand) Run(ctx context.Context) error {
	help := `
	imgproc — concurrent image processing CLI
	USAGE:
	imgproc --input <path> --output <path> [flags]

	FLAGS:
	--input     Input image or directory (required)
	--output    Output directory (required)
	--workers   Number of concurrent workers (default: total number of user OS cores)
	-h, --help  Show this help message

	EXAMPLES:
	imgproc --input ./images --output ./out
	imgproc --input photo.jpg --output ./out --workers 8
	`
	fmt.Fprintln(os.Stdout, help)
	return nil
}

type ProcessConfig struct {
	InputDir    string
	OutputDir   string
	ResizeWidth int
	Format      string
	Quality     int
	Workers     int
	Watermark   bool
	StripEXIF   bool
	MaxInflight int
}

var ErrInvalidInput = errors.New("Input dir or outputflag cannot be empty")

func (c ProcessConfig) Validate() error {
	if c.InputDir == "" || c.OutputDir == "" {
		return ErrInvalidInput
	}
	return nil
}

type ProcessCommand struct {
	cfg ProcessConfig
}

func NewProcessCommand(cfg ProcessConfig) *ProcessCommand {
	return &ProcessCommand{cfg: cfg}
}

func (p *ProcessCommand) Run(ctx context.Context) error {
	jobs := make(chan pipeline.ImageJob)

	metrics := pipeline.NewMetrics()
	metrics.Start()

	pool := pipeline.NewWorkerPool(
		p.cfg.Workers,
		p.cfg.MaxInflight,
		metrics,
	)

	var wg sync.WaitGroup

	pool.Start(ctx, jobs, &wg)

	pool.StartProgress(ctx)

	totalJobs, err := ProduceJobs(ctx, p.cfg, jobs)
	if err != nil {
		close(jobs)
		wg.Wait()
		return err
	}

	atomic.StoreInt64(&metrics.Total, int64(totalJobs))

	close(jobs)

	wg.Wait()

	fmt.Println()
	metrics.End()
	metrics.Print()

	return nil
}
