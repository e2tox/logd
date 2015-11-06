package main

import (
	"os"
	"flag"
	"log"
	"net/http"
	"text/template"
	"github.com/hpcloud/tail"
)


var addr = flag.String("addr", ":8081", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))


func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}


func main() {
	log.SetOutput(os.Stdout)
	log.Printf("Logd started")
	flag.Parse()
	go h.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	go watch()
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func watch() {

	t, err := tail.TailFile("/usr/share/tomcat/logs/catalina.out", tail.Config{Follow: true})
	if err != nil {
		log.Fatal(err)
	}
	for line := range t.Lines {
		h.notify(line.Text)
	}

	//	watcher, err := fsnotify.NewWatcher()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	defer watcher.Close()
	//
	//	done := make(chan bool)
	//	go func() {
	//		for {
	//			select {
	//			case event := <-watcher.Events:
	//				log.Println("event:", event)
	//				if event.Op&fsnotify.Write == fsnotify.Write {
	//					log.Println("modified file:", event.Name)
	//				}
	//			case err := <-watcher.Errors:
	//				log.Println("error:", err)
	//			}
	//		}
	//	}()
	//
	//	err = watcher.Add("/usr/share/tomcat/logs/catalina.out")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	<-done
}