package app

import (
	"bytes"
	"context"
	"fmt"

	"github.com/lupppig/imgproc/internal/cli"
)

func Run(ctx context.Context, commands []string) error {
	var Errbuff = new(bytes.Buffer)
	cmd, err := cli.Parse(commands, Errbuff)

	if Errbuff.Len() > 0 {
		return fmt.Errorf("imgproc: %v", Errbuff.String())
	}

	if err != nil {
		return err
	}

	return cmd.Run(ctx)
}
