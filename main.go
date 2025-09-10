package main

import (
	"log/slog"
	"os"

	"github.com/khuong02/kd-gen/cmd"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stderr, nil))
	slog.SetDefault(logger)
	if err := cmd.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
