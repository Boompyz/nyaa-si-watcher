package torrentoptions

import (
	"fmt"
	"os/exec"
	"sort"
	"strings"

	"github.com/Boompyz/nyaa-si-watcher/common"
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
	watching := common.GetLines(confDir + "/watching")
	resolved := common.GetLines(confDir + "/resolved")

	sort.Strings(resolved)

	return &ContentHandler{watching, resolved, confDir}
}

// ResetResolved clears the resolved history.
func (c *ContentHandler) ResetResolved() {
	c.resolved = make([]string, 0)
	common.WriteLines(c.confDir+"/resolved", make([]string, 0))
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

	err := common.WriteLines(c.confDir+"/resolved", c.resolved)
	return err
}

func addTorrent(link string) {
	fmt.Println("Adding: " + link)
	err := exec.Command("deluge-console", "add", link).Run()
	if err != nil {
		panic(err.Error())
	}
}

func sortedContains(arr []string, x string) bool {
	idx := sort.SearchStrings(arr, x)
	if idx == len(arr) {
		return false
	}
	return arr[idx] == x
}
