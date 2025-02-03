package file

import (
	"os"
	"slices"
	"strings"
)

const (
	lineSeparator = "\n"
)

func GetLinesFromFile(path string) ([]string, error) {
	// Opens a file by given path and returns its content as a slice of non-empty strings.
	f, err := os.ReadFile(path)
	if err != nil {
		return []string{}, err
	}

	lines := strings.Split(string(f), lineSeparator)
	lines = deleteEmptyLines(lines)

	return lines, nil
}

func deleteEmptyLines(lines []string) []string {
	// Accepts a slice of strings and returns the slice with non-empty strings.
	// Strings containing whitespaces, tabs, or end of line symbols are treated as empty.
	return slices.DeleteFunc(lines,
		func(line string) bool {
			line = strings.TrimSpace(line)
			return len(line) == 0
		})

}
