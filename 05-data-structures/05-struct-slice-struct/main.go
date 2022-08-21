package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

type car struct {
	Manufacturer string
	Model        string
	Doors        int
}

//type items struct {
//	Wisdom    []sage
//	Transport []car
//}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	jesus := sage{
		Name:  "Jesus",
		Motto: "Love all",
	}

	ford := car{
		Manufacturer: "Ford",
		Model:        "F150",
		Doors:        4,
	}

	toyota := car{
		Manufacturer: "Toyota",
		Model:        "Fit",
		Doors:        4,
	}

	sages := []sage{buddha, jesus}
	cars := []car{ford, toyota}

	//	data := items{sages, cars}
	data := struct {
		Wisdom    []sage
		Transport []car
	}{
		sages,
		cars,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
