package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="../doggy.jpeg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("../doggy.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}
func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/doggy.jpeg", dogPic)
	http.ListenAndServe(":8080", nil)
}
