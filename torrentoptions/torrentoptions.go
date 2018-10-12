package torrentoptions

import (
	"encoding/xml"
	"strings"

	"github.com/antchfx/xquery/xml"
)

// TorrentOption Represents a torrent option to download
type TorrentOption struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Size  string `xml:"size"`
}

// GetId returns the id of the torrent -> https://nyaa.si/download/<id-here>.torrent
func (t TorrentOption) GetId() string {
	parts := strings.Split(t.Link, "/")
	return strings.Split(parts[len(parts)-1], ".")[0]
}

// GetAllOptions returns all torrents in the RSS feed for HorribleSubs 720p
func GetAllOptions() ([]TorrentOption, error) {
	doc, err := xmlquery.LoadURL("https://nyaa.si/?page=rss&q=720p&c=0_0&f=0&u=HorribleSubs")
	if err != nil {
		return make([]TorrentOption, 0), err
	}

	items := xmlquery.Find(doc, "/rss/channel/item")
	torrentOptions := make([]TorrentOption, 0, len(items))

	for _, item := range items {
		var torrentOption TorrentOption
		xml.Unmarshal([]byte(item.OutputXML(true)), &torrentOption)
		torrentOptions = append(torrentOptions, torrentOption)
	}

	return torrentOptions, nil
}
