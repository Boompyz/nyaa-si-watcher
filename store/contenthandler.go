package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

// Watch represents a single query that is being watched
// and the folder the torrents should go in.
type Watch struct {
	Query  string `json:"query"`
	Folder string `json:"folder"`
	User   string `json:"user"`
}

// ContentHandler decides which files are already
// retrieved and which shouldn't be at all.
type ContentHandler struct {
	Watching  []Watch                  `json:"watching"`
	Resolved  map[string]TorrentOption `json:"resolved"`
	Announcer MailAnnouncer            `json:"mailannouncer"`

	confDir string
}

// NewContentHandler creates a new instance loaded with the
// configs in the specified folder.
func NewContentHandler(confDir string) *ContentHandler {

	filename := confDir + "/store.json"
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		// If the file doesn't exist
		_, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		var ch = &ContentHandler{make([]Watch, 0), make(map[string]TorrentOption), *NewMailAnnouncer(), confDir}
		ch.Save()
		return ch
	}

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var ret ContentHandler
	if err = json.Unmarshal(bytes, &ret); err != nil {
		panic(err)
	}
	ret.confDir = confDir

	return &ret
}

// Save the state to the store file
func (c *ContentHandler) Save() {
	encoded, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		panic(err)
	}
	//fmt.Println("Saving to file: " + string(encoded))
	if err := ioutil.WriteFile(c.confDir+"/store.json", encoded, os.ModePerm); err != nil {
		panic(err)
	}
}

// ResetResolved clears the resolved history.
func (c *ContentHandler) ResetResolved() {
	c.Resolved = make(map[string]TorrentOption)
	c.Save()
}

// GetNew gets all new required torrents and downloads them.
// Returns the newly-added torrents.
func (c *ContentHandler) GetNew() []TorrentOption {
	options := make([]TorrentOption, 0)
	for _, watch := range c.Watching {
		option, _ := GetAllOptionsWithQuery(watch.Query, watch.User)
		option = c.filterResolved(option)
		options = append(options, option...)
		c.get(option, watch.Folder)
	}
	c.Announcer.Announce(options)
	return options
}

// GetNewQuery adds a bunch of torrents without adding them to the watching
// Disregards resolved and doesn't add them there
func (c *ContentHandler) GetNewQuery(query string, folder string, user string) []TorrentOption {
	toGet, err := GetAllOptionsWithQuery(query, user)
	if err != nil {
		fmt.Print(err)
	}
	for _, option := range toGet {
		addTorrent(option.Link, folder)
	}
	return toGet
}

// AddNewWatch adds the new watch
func (c *ContentHandler) AddNewWatch(watch string, folder string, user string) {
	c.Watching = append(c.Watching, Watch{watch, folder, user})
	c.Save()
}

// RemoveWatch removes a watch
func (c *ContentHandler) RemoveWatch(watch string) {
	index := -1
	for idx, val := range c.Watching {
		if val.Query == watch {
			index = idx
			break
		}
	}
	if index != -1 {
		c.Watching = append(c.Watching[:index], c.Watching[index+1:]...)
	}
	c.Save()
}

// filterResolved filters the options to include only the ones that are not resolved
func (c *ContentHandler) filterResolved(options []TorrentOption) []TorrentOption {
	filteredOptions := make([]TorrentOption, 0)

	for _, option := range options {
		_, contained := c.Resolved[option.GetID()]
		if !contained {
			filteredOptions = append(filteredOptions, option)
		}
	}

	return filteredOptions
}

// Get the files specified. Overwrites resolved in the conf directory.
func (c *ContentHandler) get(options []TorrentOption, folder string) {
	for _, option := range options {
		addTorrent(option.Link, folder)
		c.Resolved[option.GetID()] = option
	}
	c.Save()
}

func addTorrent(link string, folder string) {
	fmt.Println("Adding: " + link)
	cmd := exec.Command("curl", "-O", link)
	cmd.Dir = folder
	err := cmd.Run()
	if err != nil {
		panic(err.Error())
	}
}
