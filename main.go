package main

import (
	"flag"

	"github.com/Boompyz/nyaa-si-watcher/webui"
)

func main() {

	var confDir = flag.String("confDir", "/var/lib/nyaa-si-watcher", "The directory to look for watching.conf and resolved.conf")
	//var resetResolved = flag.Bool("r", false, "Set this to reset already resolved files. Default is false.")

	flag.Parse()

	webui.Listen(*confDir)
}
