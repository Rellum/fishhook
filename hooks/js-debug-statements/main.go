package main

import (
	"fmt"
	"github.com/Rellum/fishhook/pkg/forbidden"
	"io"
	"os"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	return forbidden.CheckFiles(args, []string{
		"nocommit",
	})
}
