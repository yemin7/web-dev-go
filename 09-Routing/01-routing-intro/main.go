package main

import (
	"io"
	"net/http"
)

type pet int

func (p pet) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/cat":
		io.WriteString(w, "kitty kitty")
	case "/dog":
		io.WriteString(w, "doggy doggy")
	}
}

func main() {
	var p pet
	http.ListenAndServe(":8080", p)
}
