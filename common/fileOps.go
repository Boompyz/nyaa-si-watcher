package common

import (
	"bufio"
	"os"
)

// GetLines reads all lines from a file.
func GetLines(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Couldn't read file: " + fileName)
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// WriteLines writes a bunch of lines to a file
func WriteLines(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	file.Close()
	return nil
}
