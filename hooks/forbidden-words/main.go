package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"os"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	argsY, err := yaml.Marshal(args)
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Args: %+v\n", string(argsY))

	envY, err := yaml.Marshal(os.Environ())
	if err != nil {
		return err
	}
	fmt.Fprintf(stdout, "Env: %+v\n", string(envY))

	return nil
}
