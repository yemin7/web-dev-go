package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var tpl *template.Template

var funcmap = template.FuncMap{
	"fdouble": double,
	"fsquare": square,
	"fsqroot": sqroot,
}

func init() {
	tpl = template.Must(template.New("").Funcs(funcmap).ParseFiles("tpl.gohtml"))
}

func double(f int) int {
	return int(f + f)
}

func square(f int) float64 {
	return math.Pow(float64(f), 2)
}

func sqroot(f float64) float64 {
	return math.Sqrt(f)
}

func main() {
	num := 3
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", num)
	if err != nil {
		log.Fatalln(err)
	}
}
