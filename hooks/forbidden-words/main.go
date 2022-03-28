package main

import (
	"errors"
	"fmt"
	"github.com/Rellum/fishhook/pkg/forbidden"
	"io"
	"os"
	"strings"
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
	if len(args) < 3 { // /path/to/tool "forbidden word list" dir/file.go
		return errors.New("no forbidden words passed")
	}

	var forbiddenWords []string
	if len(args) > 1 {
		forbiddenWords = strings.Split(args[1], " ")
		for i := range forbiddenWords {
			forbiddenWords[i] = strings.Trim(forbiddenWords[i], "\"'`")
		}
	}

	return forbidden.CheckFiles(args, forbiddenWords)
}
