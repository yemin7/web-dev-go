package main

import (
	"io"
	"log"
	"net/http"
	"text/template"
)

func foo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func dog(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("dog.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	tpl.ExecuteTemplate(w, "dog.gohtml", nil)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "doggy.jpeg")
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/doggy.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}
