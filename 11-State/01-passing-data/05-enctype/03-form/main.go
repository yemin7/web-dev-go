package main

import (
	"log"
	"net/http"
	"text/template"
)

var tpl *template.Template

type person struct {
	FirstName, LastName string
	Subscribed          bool
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	bs := make([]byte, req.ContentLength)
	req.Body.Read(bs)
	body := string(bs)

	err := tpl.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
