package main

import (
	"github.com/google/uuid"
	"net/http"
	"text/template"
)

type user struct {
	UserName, First, Last string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //userID, user
var dbSessions = map[string]string{} //sessionID, userID

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, cookie)
	}

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")
		u = user{un, f, l}
		dbSessions[cookie.Value] = un
		dbUsers[un] = u
	}
	tpl.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.gohtml", u)
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}
