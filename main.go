package main

import (
	"fmt"
	"os"

	"github.com/khuong02/kd-gen/cmd"
)

func main() {
	if err := cmd.Run(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "Exit: %v\n", err)
		os.Exit(1)
	}
}
