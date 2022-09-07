package main

import (
	"io"
	"net/http"
)

func ServeDog(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy Doggy")
}

func ServeCat(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Kitty Kitty")
}

func main() {
	http.HandleFunc("/cat/", ServeCat)
	http.HandleFunc("/dog", ServeDog)
	http.ListenAndServe(":8080", nil)
}
