package main

import (
	"os"

	"github.com/khuong02/kd-gen/cmd"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		os.Exit(1)
	}
}
