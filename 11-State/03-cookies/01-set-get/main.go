package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "my-cookie",
		Value: "some value",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN - CHECK BROWSER")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("my-cookie")
	if err != nil {
		http.Error(w, "Error pop naw", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "YOUR COOKIE", cookie)
}
