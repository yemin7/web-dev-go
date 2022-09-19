package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func dog(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("template didn't execute: ", err)
	}
}

func main() {
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/pics", fs)
	http.HandleFunc("/", dog)
	err := http.ListenAndServe(":8080", fs)
	if err != nil {
		return
	}
}
