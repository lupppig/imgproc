package cli

import (
	"context"
	"fmt"
	"os"
)

type HelpExecutable struct {
	Err error
}

func (h *HelpExecutable) Run(ctx context.Context) {
	help := `
	imgproc — concurrent image processing CLI
	USAGE:
	imgproc --input <path> --output <path> [flags]

	FLAGS:
	--input     Input image or directory (required)
	--output    Output directory (required)
	--workers   Number of concurrent workers (default: 4)
	-h, --help  Show this help message

	EXAMPLES:
	imgproc --input ./images --output ./out
	imgproc --input photo.jpg --output ./out --workers 8
	`
	fmt.Fprintln(os.Stdout, help)
}

func (h *HelpExecutable) Error() error {
	return nil
}

type ProcessCommand struct {
	InputDir  string
	OutputDir string
	Format    string

	Workers     int
	ResizeWidth int
	Quality     int
	MaxInflight int

	Err error
}

func (p *ProcessCommand) Run(ctx context.Context) {
	fmt.Printf("welp what a command to run !!!!!!!")
}

func (p *ProcessCommand) Error() error {
	return p.Err
}
