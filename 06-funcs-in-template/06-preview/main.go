package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	err := tpl.ExecuteTemplate(writer, "index.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
