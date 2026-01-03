package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lupppig/imgproc/internal/app"
	"github.com/lupppig/imgproc/internal/commands"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := app.Run(ctx, os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "imgproc:", err)
		os.Exit(commands.EXIT_FAILURE)
	}

}
