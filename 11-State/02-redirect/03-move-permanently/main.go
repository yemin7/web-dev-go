package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/foobar", foobar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo: ", req.Method)
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar: ", req.Method)
	http.Redirect(w, req, "/", http.StatusMovedPermanently)
}

func foobar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foobar: ", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
