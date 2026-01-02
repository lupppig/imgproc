package cli

import (
	"flag"
	"io"
	"runtime"

	"github.com/lupppig/imgproc/internal/app"
	"github.com/lupppig/imgproc/internal/config"
)

func Parse(args []string, writer io.Writer) app.Flag {
	fs := flag.NewFlagSet("imgproc", flag.ContinueOnError)
	fs.SetOutput(writer)

	cfg := new(config.Config)

	var help bool
	flag.StringVar(&cfg.InputDir, "input", "", "Input directory")
	flag.StringVar(&cfg.OutputDir, "output", "", "Output directory")
	flag.IntVar(&cfg.ResizeWidth, "resize", 0, "Resize width")
	flag.StringVar(&cfg.Format, "format", "jpeg", "Output format")
	flag.IntVar(&cfg.Quality, "quality", 80, "JPEG quality")
	flag.IntVar(&cfg.Workers, "workers", runtime.NumCPU(), "Processing workers")
	flag.IntVar(&cfg.MaxInflight, "max-inflight", 50, "Max images in pipeline")
	flag.BoolVar(&help, "help", false, "print command help")

	flag.Parse()

	if help {
		return &HelpExecutable{}
	}

	if err := Validate(cfg); err != nil {
		return &ProcessCommand{Err: err}
	}

	p := &ProcessCommand{
		InputDir:    cfg.InputDir,
		Format:      cfg.Format,
		ResizeWidth: cfg.ResizeWidth,
		Quality:     cfg.Quality,
		MaxInflight: cfg.MaxInflight,
	}
	return p
}
