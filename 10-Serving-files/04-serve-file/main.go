package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="../doggy.jpeg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "../doggy.jpeg")
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/doggy.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}
