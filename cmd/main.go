package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lupppig/imgproc/internal/app"
	"github.com/lupppig/imgproc/internal/cli"
	"github.com/lupppig/imgproc/internal/config"
)

func main() {
	cfg := cli.Parse()

	if err := cli.Validate(cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(config.EXIT_FAILURE)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx, cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(config.EXIT_FAILURE)
	}

}
