package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	score := struct {
		Score1, Score2 int
	}{7, 9}

	err := tpl.Execute(os.Stdout, score)
	if err != nil {
		log.Fatalln(err)
	}
}
