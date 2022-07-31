package main

import (
	"net/http"
	"text/template"

	"github.com/sleong110/go-web-programming/ch2_chit_chat_app/data"
)

func index(w http.ResponseWriter, r *http.Request) {
	files := []string{"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html"}
	templates := template.Must(template.ParseFiles(files))
	threads, err := data.Threads()
	if err == nil {
		templates.ExecuteTemplate(w, "layout", threads)
	}
}

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	http.ListenAndServe(":8080", mux)
}
