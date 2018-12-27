package webui

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Boompyz/nyaa-si-watcher/store"
)

var confDir string
var contentHandler *store.ContentHandler
var tmpl *template.Template

func update() {
	contentHandler.GetNew()
}

type pageData struct {
	ContentHandler store.ContentHandler
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

	contentHandler.GetNewQuery(request)
	fmt.Println(request)
}

func addWatch(w http.ResponseWriter, r *http.Request) {
	requestByte, _ := ioutil.ReadAll(r.Body)
	request := string(requestByte)

	contentHandler.AddNewWatch(request)
	fmt.Println(request)
}

func ignore(w http.ResponseWriter, r *http.Request) {

}

func Listen(_confDir string, _port int, interval int) {
	tmpl = template.Must(template.ParseFiles("webui/template.html"))

	confDir = _confDir
	contentHandler = store.NewContentHandler(confDir)

	go func() {
		for range time.NewTicker(time.Second * time.Duration(interval)).C {
			update()
		}
	}()

	fs := http.FileServer(http.Dir("webui"))

	http.Handle("/script.js", fs)
	http.HandleFunc("/", handler)
	http.HandleFunc("/request", requests)
	http.HandleFunc("/addwatch", addWatch)
	http.HandleFunc("/favicon.ico", ignore)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(_port), nil))
}
