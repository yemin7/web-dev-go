package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func root(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func me(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "me.gohtml", nil)
}

func main() {
	http.Handle("/", http.HandlerFunc(root))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))
	http.ListenAndServe(":8080", nil)
}
