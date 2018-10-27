package webui

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

	get(options)
}

func get(options []torrentoptions.TorrentOption) {

	err := contentHandler.Get(options)
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

func requests(w http.ResponseWriter, r *http.Request) {
	requestByte, _ := ioutil.ReadAll(r.Body)
	request := string(requestByte)
	strings.Replace(request, " ", "+", -1)

	options, _ := torrentoptions.GetAllOptionsWithQuery(request)
	fmt.Println(options)
	get(options)
	fmt.Println(request)
}

func ignore(w http.ResponseWriter, r *http.Request) {

}

func Listen(_confDir string, _port int, interval int) {
	tmpl = template.Must(template.ParseFiles("webui/template.html"))

	go func() {
		for range time.NewTicker(time.Second * time.Duration(interval)).C {
			update()
		}
	}()

	confDir = _confDir
	contentHandler = torrentoptions.NewContentHandler(confDir)

	fs := http.FileServer(http.Dir("webui"))

	http.Handle("/script.js", fs)
	http.HandleFunc("/", handler)
	http.HandleFunc("/request", requests)
	http.HandleFunc("/favicon.ico", ignore)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(_port), nil))
}
