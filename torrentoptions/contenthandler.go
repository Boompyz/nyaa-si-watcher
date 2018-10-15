package torrentoptions

import (
	"bufio"
	"os"
	"os/exec"
	"sort"
	"strings"
)

// ContentHandler decides which files are already
// retrieved and which shouldn't be at all.
type ContentHandler struct {
	watching []string
	resolved []string

	confDir string
}

// NewContentHandler creates a new instance loaded with the
// configs in the specified folder.
func NewContentHandler(confDir string) *ContentHandler {
	watching := getLines(confDir + "/watching")
	resolved := getLines(confDir + "/resolved")

	sort.Strings(resolved)

	return &ContentHandler{watching, resolved, confDir}
}

// Filter the options to include only the ones that contain any of
// watching and are not in resolved.
func (c *ContentHandler) Filter(options []TorrentOption) []TorrentOption {

	filteredOptions := make([]TorrentOption, 0)

	for _, option := range options {
		for _, watchedTitle := range c.watching {
			if strings.Contains(option.Title, watchedTitle) && !sortedContains(c.resolved, option.GetID()) {
				filteredOptions = append(filteredOptions, option)
				break
			}
		}
	}

	return filteredOptions
}

// Get the files specified. Overwrites resolved in the conf directory.
func (c *ContentHandler) Get(options []TorrentOption) error {
	newlyResolved := make([]string, 0, len(options))
	for _, option := range options {
		addTorrent(option.Link)
		newlyResolved = append(newlyResolved, option.GetID())
	}
	c.resolved = append(c.resolved, newlyResolved...)

	err := writeLines(c.confDir+"/resolved", c.resolved)
	return err
}

func addTorrent(link string) {
	err := exec.Command("deluge-console", "add", link).Start()
	if err != nil {
		panic(err.Error())
	}
}

func getLines(fileName string) []string {
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

func writeLines(fileName string, lines []string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}

	for _, line := range lines {
		_, err := file.WriteString(line)
		if err != nil {
			return err
		}
	}

	file.Close()
	return nil
}

func sortedContains(arr []string, x string) bool {
	idx := sort.SearchStrings(arr, x)
	if idx == len(arr) {
		return false
	}
	return arr[idx] == x
}
