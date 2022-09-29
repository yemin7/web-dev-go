package main

import "net/http"

func getUser(req *http.Request) user {
	var u user

	cookie, err := req.Cookie("session")
	if err != nil {
		return u
	}

	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}
