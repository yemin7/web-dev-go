// POST => body
// GET => url
package main

import (
	"io"
	"log"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := io.WriteString(w, `
	<form method="GET">
	 <input type="text" name="q">
	 <input type="submit">
	</form>
	<br>`+v)
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
