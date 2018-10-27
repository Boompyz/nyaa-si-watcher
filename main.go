package main

import (
	"flag"

	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
	"github.com/Boompyz/nyaa-si-watcher/webui"
)

func main() {

	var confDir = flag.String("confDir", "/var/lib/nyaa-si-watcher", "The directory to look for watching.conf and resolved.conf")
	var port = flag.Int("port", 21037, "Port to listen on")
	var interval = flag.Int("interval", 30, "Time in seconds between checks")
	var defaultResolution = flag.String("res", "720p", "Default resolution to look for files in (example: 720p or 1080p)")
	//var resetResolved = flag.Bool("r", false, "Set this to reset already resolved files. Default is false.")

	flag.Parse()

	torrentoptions.DefaultResolution = *defaultResolution

	webui.Listen(*confDir, *port, *interval)
}
