package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/turbobytes/certmon/pkg/certmon"
)

var (
	results  certmon.Results
	cfgfile  = flag.String("config", "config.yaml", "Path to config file. Autoreloads if changed")
	addr     = flag.String("listen", ":8081", "Address to listen on")
	htmlFile = flag.String("ui", "assets/index.html", "path to index.html")
)

func init() {
	flag.Parse()
}

//TODO: Prometheus things

func handleResult(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(results)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("last-modified", results.Timestamp.Format(http.TimeFormat))
	w.Header().Set("Cache-control", "must-revalidate")
	//Check if modified since
	ims := r.Header.Get("if-modified-since")
	if ims != "" {
		t, err := time.Parse(http.TimeFormat, ims)
		if err == nil {
			if !results.Timestamp.Truncate(time.Second).After(t) {
				w.WriteHeader(http.StatusNotModified)
				return
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func runtest(config certmon.Config) {
	log.Println("Running tests")
	results = config.Run()
	log.Println("Done")
}

func main() {
	//Load from yaml or smthn and attach inotify
	config, err := certmon.LoadConfig(*cfgfile)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/results/", handleResult)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, *htmlFile)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {})
	go func() {
		log.Fatal(http.ListenAndServe(*addr, nil))
	}()
	log.Println("Listening on " + *addr)

	//Create a watcher on config file
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	watcher.Add(*cfgfile)
	//Run once on start
	runtest(config)
	for {
		select {
		case <-time.After(config.LoopDuration):
		case <-watcher.Events:
			log.Println("Config file changed")
			cfg, err := certmon.LoadConfig(*cfgfile)
			if err != nil {
				log.Println("Error loading config file: " + err.Error())
			} else {
				config = cfg
			}
		}
		runtest(config)
	}
}
