package main

import (
	"io"
	"net/http"
)

type dog int
type cat int

func (d dog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy doggy")
}

func (c cat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty kitty")
}

func main() {
	var c cat
	var d dog

	mux := http.NewServeMux()
	mux.Handle("/cat", c)
	mux.Handle("/dog/", d)

	//DefaultServeMux
	//http.Handle("/cat", c)
	//http.Handle("/dog/", d)
	http.ListenAndServe(":8080", mux)
}
