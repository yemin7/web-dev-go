package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	str := []string{"A", "B", "C", "D", "E"}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", str)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-var.gohtml", str)
	if err != nil {
		log.Fatalln(err)
	}
}
