package store

import (
	"encoding/xml"
	"net/url"
	"strings"

	xmlquery "github.com/antchfx/xquery/xml"
)

// TorrentOption Represents a torrent option to download
type TorrentOption struct {
	Title string `xml:"title" json:"title"`
	Link  string `xml:"link" json:"link"`
	Size  string `xml:"size" json:"size"`
}

// GetID returns the id of the torrent -> https://nyaa.si/download/<id-here>.torrent
func (t TorrentOption) GetID() string {
	parts := strings.Split(t.Link, "/")
	return strings.Split(parts[len(parts)-1], ".")[0]
}

// GetAllOptionsWithQuery returns all torrents in the RSS feed for HorribleSubs with the given query
// The function handles url character escape
func GetAllOptionsWithQuery(searchString string) ([]TorrentOption, error) {
	searchString = url.QueryEscape(searchString)
	doc, err := xmlquery.LoadURL("https://nyaa.si/?page=rss&q=" + searchString + "&c=0_0&f=0&u=HorribleSubs")
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
