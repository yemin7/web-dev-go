package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (c hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("My-Key", "This is my key.")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code in the function.</h1>")
}

func main() {
	var c hotdog
	http.ListenAndServe(":8080", c)
}
