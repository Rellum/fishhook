package forbidden

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckFiles(files, forbidden []string) error {
	for i := range files {
		if err := checkFile(files[i], forbidden); err != nil {
			return err
		}
	}

	return nil
}

func checkFile(name string, forbidden []string) error {
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
