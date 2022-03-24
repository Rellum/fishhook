package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	// exitFail is the exit code if the program fails.
	exitFail = 1
)

var forbidden = []string{
	"nocommit",
	"console.log(",
	"debugger",
}

func main() {
	if err := run(os.Args, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(exitFail)
	}
}

func run(args []string, stdout io.Writer) error {
	for i := range args[1:] {
		checkFile(args[i])
	}

	return nil
}

func checkFile(name string) error {
	file, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("read file error: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var i int
	for scanner.Scan() {
		i++
		text := scanner.Text()

		for j := range forbidden {
			idx := strings.Index(text, forbidden[j])
			if idx >= 0 {
				return fmt.Errorf("forbidden string '%s' found in %s:%v col %v", forbidden[j], name, i, idx)
			}
		}
	}

	return nil
}
