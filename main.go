package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
)

func addTorrent(link string) {
	err := exec.Command("deluge-console", "add", link).Start()
	if err != nil {
		panic(err.Error())
	}
}

func getWatches(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		panic("Couldn't read file")
	}
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	options, err := torrentoptions.GetAllOptions()
	if err != nil {
		panic(err.Error())
	}

	watched := getWatches("watch/watching.txt")
	for _, option := range options {
		get := false
		for _, watchedTitle := range watched {
			if strings.Contains(option.Title, watchedTitle) {
				get = true
				break
			}
		}

		if get {
			addTorrent(option.Link)
		}
	}
}
