package main

import (
	"net/http"
	"text/template"
)

type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	page := Page{"ajax test.", 1}
	tmpl, err := template.ParseFiles("main.html") // ParseFilesを使う
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(w, page)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
