package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type it struct {
	Type     string
	Language string
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}
func main() {

	it_go := it{"Programming", "Golang"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", it_go)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-var.gohtml", it_go)
	if err != nil {
		log.Fatalln(err)
	}
}
