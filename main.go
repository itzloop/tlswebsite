package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path"
)

// i will never forgive my self for writing code like this but clocks ticking :)

type Server struct {
}

func handleError(w http.ResponseWriter) error {
	t, err := template.ParseFiles(path.Join("./", "template", "error.gohtml"))
	if err != nil {
		return err
	}

	if err = t.Execute(w, nil); err != nil {
		return err
	}

	return nil
}

func handler(r http.ResponseWriter, _ *http.Request) {
	indexPath := path.Join("./", "static", "index.html")
	f, err := os.ReadFile(indexPath)
	if err != nil {
		log.Printf("handler: failed to open index.html: %s", err.Error())
		if err = handleError(r); err != nil {
			log.Printf("handler: failed to handle error: %s", err.Error())
		}
		return
	}

	_, err = r.Write(f)
	if err != nil {
		log.Printf("handler: failed to write response: %s", err.Error())
	}
}

func main() {
	addr, exists := os.LookupEnv("ADDR")
	if !exists || addr == "" {
		addr = ":8080"
	}

	fs := http.FileServer(http.Dir(path.Join("./", "static/")))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", handler)

	log.Printf("server is starting at %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
