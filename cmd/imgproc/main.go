package main

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/lupppig/imgproc/internal/cli"
	"github.com/lupppig/imgproc/internal/config"
)

func main() {
	var buff = new(bytes.Buffer)

	var commands = make([]string, len(os.Args[1:]))
	for _, c := range os.Args[1:] {
		cmd := strings.Split(c, "=")
		commands = append(commands, cmd[0])
	}

	cmd := cli.Parse(commands, buff)

	if buff.Len() != 0 {
		fmt.Fprintln(os.Stderr, "imgproc: "+buff.String())
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := cmd.Error(); err != nil {
		fmt.Fprintln(os.Stderr, "imgproc: "+err.Error())
		os.Exit(config.EXIT_FAILURE)
	}

	cmd.Run(ctx)
}
