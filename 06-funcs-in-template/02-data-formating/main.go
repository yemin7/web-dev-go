package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

var funcmap = template.FuncMap{
	"fdateMDY": tformat,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("tpl.gohtml"))
}

func tformat(t time.Time) string {
	return t.Format("01-02-2006")
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
