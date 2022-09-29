package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, `<h1><a href="/set">set a cookie</a></h1>`)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "cookie-1",
		Value: "0",
		Path:  "/",
	})
	fmt.Fprintln(w, `<h1><a href="/set">read</a></h1>`)
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie-1")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	fmt.Fprintf(w, `<h1>Your Cookie:<br>%v</h1><h1><a href="/expire">expire</a></h1>`, cookie)
}

func expire(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("cookie-1")
	if err != nil {
		http.Redirect(w, req, "/set", http.StatusSeeOther)
		return
	}
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, req, "/", http.StatusSeeOther)
}
