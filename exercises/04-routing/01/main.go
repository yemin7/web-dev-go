package main

import (
	"io"
	"net/http"
)

func root(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "")
}

func dog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy Doggy")
}

func me(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "It is me")
}

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
