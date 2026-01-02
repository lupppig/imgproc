package cli

import (
	"flag"
	"runtime"

	"github.com/lupppig/imgproc/internal/config"
)

func Parse() *config.Config {
	cfg := new(config.Config)

	flag.StringVar(&cfg.InputDir, "input", "", "Input directory")
	flag.StringVar(&cfg.OutputDir, "output", "", "Output directory")
	flag.IntVar(&cfg.ResizeWidth, "resize", 0, "Resize width")
	flag.StringVar(&cfg.Format, "format", "jpeg", "Output format")
	flag.IntVar(&cfg.Quality, "quality", 80, "JPEG quality")
	flag.IntVar(&cfg.Workers, "workers", runtime.NumCPU(), "Processing workers")
	flag.IntVar(&cfg.MaxInflight, "max-inflight", 50, "Max images in pipeline")

	flag.Parse()
	return cfg
}
