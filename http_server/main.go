package main

import (
	"html/template"
	"net/http"
	"log"
)

func check(err error) {
	if err!= nil {
		log.Fatal(err)
	}
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("index.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func main() {
	http.HandleFunc("/", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
