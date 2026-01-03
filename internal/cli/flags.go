package cli

import (
	"flag"
	"io"
	"runtime"

	"github.com/lupppig/imgproc/internal/commands"
)

func Parse(args []string, out io.Writer) (commands.Command, error) {
	fs := flag.NewFlagSet("imgproc", flag.ContinueOnError)
	fs.SetOutput(out)

	cfg := commands.ProcessConfig{}
	help := false

	fs.StringVar(&cfg.InputDir, "input", "", "Input directory")
	fs.StringVar(&cfg.OutputDir, "output", "", "Output directory")
	fs.IntVar(&cfg.ResizeWidth, "resize", 0, "Resize width")
	fs.StringVar(&cfg.Format, "format", "jpeg", "Output format")
	fs.IntVar(&cfg.Quality, "quality", 80, "Image quality")
	fs.IntVar(&cfg.Workers, "workers", runtime.NumCPU(), "Workers")
	fs.IntVar(&cfg.MaxInflight, "max-inflight", 50, "Inflight limit")
	fs.BoolVar(&cfg.Watermark, "watermark", false, "remove watermake from image background")
	fs.BoolVar(&cfg.StripEXIF, "strip-exif", false, "strip exif from image")
	fs.BoolVar(&help, "help", false, "Show help")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if help {
		return &commands.HelpCommand{}, nil
	}

	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return commands.NewProcessCommand(cfg), nil
}
