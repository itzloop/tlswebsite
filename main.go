package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

// i will never forgive my self for writing code like this but clocks ticking :)

type Server struct {
}

func handler(r http.ResponseWriter, _ *http.Request) {
	indexPath := path.Join("./", "static", "index.html")
	f, err := os.ReadFile(indexPath)
	if err != nil {
		log.Printf("handler: failed to open index.html: %s", err.Error())
	}

	_, err = r.Write(f)
	if err != nil {
		log.Printf("handler: failed to write response: %s", err.Error())
	}
}

func main() {
	addr := os.Args[1]

	fs := http.FileServer(http.Dir(path.Join("./", "static/")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)

	http.ListenAndServe(addr, nil)
}
