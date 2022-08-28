package main

import (
	"html/template"
	"log"
	"os"
)

type hotel struct {
	Name, Address, City string
	Zip                 int
}

type region struct {
	Region string
	Hotels []hotel
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	hotel_list := []region{
		region{
			Region: "Southern",
			Hotels: []hotel{
				hotel{"Hotel California", "42 Sunset Boulevard", "Los Angeles", 95612},
				hotel{"Hotel California2", "43 Sunset", "LA", 95612},
			},
		},
		region{
			Region: "Northern",
			Hotels: []hotel{
				hotel{"Hotel California", "42 Sunset Boulevard", "Los Angeles", 95612},
				hotel{"Hotel California2", "43 Sunset", "LA", 95612},
			},
		},
	}

	err := tpl.Execute(os.Stdout, hotel_list)
	if err != nil {
		log.Fatalln(err)
	}
}
