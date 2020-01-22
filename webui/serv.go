package webui

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/Boompyz/nyaa-si-watcher/store"
)

type query struct {
	Query  string `json:"query"`
	Folder string `json:"folder"`
	User   string `json:"user"`
}

var confDir string
var contentHandler *store.ContentHandler
var tmpl *template.Template

var mutex sync.Mutex

func update() {
	mutex.Lock()
	contentHandler.GetNew()
	mutex.Unlock()
}

type pageData struct {
	ContentHandler store.ContentHandler
}

func pageServe(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
	//fmt.Println(contentHandler)
	tmpl.Execute(w, pageData{
		ContentHandler: *contentHandler,
	})
}

func requests(w http.ResponseWriter, r *http.Request) {
	requestByte, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(requestByte))
	q := query{}
	if err := json.Unmarshal(requestByte, &q); err != nil {
		fmt.Println(err.Error())
	}

	contentHandler.GetNewQuery(q.Query, q.Folder, q.User)
}

func removeWatch(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	requestByte, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(requestByte))
	q := query{}
	if err := json.Unmarshal(requestByte, &q); err != nil {
		fmt.Println(err.Error())
	}
	contentHandler.RemoveWatch(q.Query)
	w.Write([]byte("Removed" + q.Query))
	mutex.Unlock()
}

func addWatch(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	requestByte, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(requestByte))
	q := query{}
	if err := json.Unmarshal(requestByte, &q); err != nil {
		fmt.Println(err.Error())
	}
	w.Write([]byte("Added"))
	contentHandler.AddNewWatch(q.Query, q.Folder, q.User)
	contentHandler.GetNew()
	mutex.Unlock()
}

func ignore(w http.ResponseWriter, r *http.Request) {
	// ignore the request
}

func addEmail(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	requestByte, _ := ioutil.ReadAll(r.Body)
	request := string(requestByte)

	contentHandler.Announcer.AddEmail(request)
	contentHandler.Save()
	w.Write([]byte("Added"))
	mutex.Unlock()
}

func removeEmail(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	requestByte, _ := ioutil.ReadAll(r.Body)
	request := string(requestByte)

	contentHandler.Announcer.RemoveEmail(request)
	contentHandler.Save()
	w.Write([]byte("Removed"))
	mutex.Unlock()
}

// Listen starts the listeners for the web interface.
func Listen(_confDir string, _port int, interval int) {
	tmpl = template.Must(template.ParseFiles("webui/template.html"))

	confDir = _confDir
	contentHandler = store.NewContentHandler(confDir)

	go func() {
		for {
			update()
			time.Sleep(time.Duration(interval) * time.Second)
		}
	}()

	fs := http.FileServer(http.Dir("webui"))

	http.Handle("/script.js", fs)
	http.HandleFunc("/", pageServe)
	http.HandleFunc("/request", requests)

	http.HandleFunc("/addwatch", addWatch)
	http.HandleFunc("/removewatch", removeWatch)

	http.HandleFunc("/addemail", addEmail)
	http.HandleFunc("/removeemail", removeEmail)

	http.HandleFunc("/favicon.ico", ignore)

	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(_port), nil))
}
