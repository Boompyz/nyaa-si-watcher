package webui

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/Boompyz/nyaa-si-watcher/announcer"
	"github.com/Boompyz/nyaa-si-watcher/torrentoptions"
)

var confDir string
var contentHandler *torrentoptions.ContentHandler
var tmpl *template.Template

func update() {

	options, err := torrentoptions.GetAllOptions()
	if err != nil {
		panic(err.Error())
	}

	options = contentHandler.Filter(options)
	for _, option := range options {
		fmt.Println("Found new: " + option.Title)
	}

	announcer := announcer.NewMailAnnouncer(confDir)
	announcer.Announce(options)

	err = contentHandler.Get(options)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Added successfuly")
}

type pageData struct {
	ContentHandler torrentoptions.ContentHandler
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	fmt.Println(contentHandler)
	tmpl.Execute(w, pageData{
		ContentHandler: *contentHandler,
	})
}

func ignore(w http.ResponseWriter, r *http.Request) {

}

func Listen(_confDir string) {
	tmpl = template.Must(template.ParseFiles("webui/template.html"))

	go func() {
		for range time.NewTicker(time.Second * 30).C {
			update()
		}
	}()

	confDir = _confDir
	contentHandler = torrentoptions.NewContentHandler(confDir)

	http.HandleFunc("/", handler)
	http.HandleFunc("/favicon.ico", ignore)
	log.Fatal(http.ListenAndServe(":80", nil))
}
