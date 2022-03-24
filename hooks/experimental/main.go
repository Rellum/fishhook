package main

import (
	"errors"
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
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout, stderr io.Writer) error {
	argsY, err := yaml.Marshal(args)
	if err != nil {
		return err
	}

	fmt.Fprint(stdout, "This is stdout\n")
	fmt.Fprint(stderr, "This is stderr\n")

	w := io.MultiWriter(stdout, os.Stderr)
	fmt.Fprintf(w, "Args: %+v\n", string(argsY))

	envY, err := yaml.Marshal(os.Environ())
	if err != nil {
		return err
	}
	fmt.Fprintf(w, "Env: %+v\n", string(envY))

	for _, arg := range args {
		if arg == "fail" {
			return errors.New("fail")
		}
	}

	return nil
}
