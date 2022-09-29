package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/users/spell", set)
	http.HandleFunc("/users/spell/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie-1")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "COOKIE ", cookie)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie-1",
		Value: "Testing 1 2 3",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie-1")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "COOKIE ", cookie)
}
