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
	str_map := map[string]string{
		"Type":     "Programming",
		"Language": "Golang",
	}
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", str_map)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "tpl-var.gohtml", str_map)
	if err != nil {
		log.Fatalln(err)
	}
}
