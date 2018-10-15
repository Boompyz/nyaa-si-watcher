package main

import (
	"flag"
	"fmt"

	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
)

func main() {

	var confDir = flag.String("confDir", "/var/lib/nyaa-si-watcher", "The directory to look for watching.conf and resolved.conf")

	contentHandler := torrentoptions.NewContentHandler(*confDir)

	options, err := torrentoptions.GetAllOptions()
	if err != nil {
		panic(err.Error())
	}

	options = contentHandler.Filter(options)
	for _, option := range options {
		fmt.Println("Found new: " + option.Title)
	}

	err = contentHandler.Get(options)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Added successfuly")
}
