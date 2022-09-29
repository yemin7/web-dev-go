package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", setone)
	http.HandleFunc("/multi", setmulti)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func setone(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("single-cookie")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "single-cookie",
			Value: "0",
			Path:  "/",
		}
	}
	count, _ := strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)
	io.WriteString(w, cookie.Value)
}

func setmulti(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie-1",
		Value: "Multi cookie 1",
		Path:  "/",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "cookie-2",
		Value: "Multi cookie 2",
		Path:  "/",
	})
	fmt.Fprintln(w, "COOKIE WRITTEN")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie1, err := req.Cookie("single-cookie")
	if err != nil {
		http.Error(w, "Error pop naw", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Single Cookie: ", cookie1)

	cookie2, err := req.Cookie("cookie-1")
	if err != nil {
		http.Error(w, "Error pop naw", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Multi Cookie1: ", cookie2)

	cookie3, err := req.Cookie("cookie-2")
	if err != nil {
		http.Error(w, "Error pop naw", http.StatusBadRequest)
	}
	fmt.Fprintln(w, "Multi Cookie2: ", cookie3)
}
