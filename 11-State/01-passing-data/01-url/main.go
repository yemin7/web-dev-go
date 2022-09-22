package main

import (
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	_, err := io.WriteString(w, "Do your search: "+v)
	if err != nil {
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
