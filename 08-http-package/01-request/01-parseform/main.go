package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type hotdog int

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func check(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	check(err)
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
