package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type user struct {
	Name   string
	Singer bool
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	user1 := user{
		Name:   "Lee Jieun",
		Singer: true,
	}

	user2 := user{
		Name:   "Han So Hee",
		Singer: false,
	}

	users := []user{user1, user2}

	err := tpl.Execute(os.Stdout, users)
	if err != nil {
		log.Fatalln(err)
	}
}
