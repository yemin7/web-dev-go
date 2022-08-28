package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type item struct {
	Name        string
	Description string
	Price       float64
}

type menu struct {
	Items []item
}

type meal struct {
	Breakfast, Lunch, Dinner menu
}

func main() {
	meals := []meal{
		meal{
			Breakfast: menu{
				Items: []item{
					item{"Bread", "cheese", 22.0},
					item{"Milk", "plain", 16.5},
				},
			},
		},
		meal{
			Lunch: menu{
				Items: []item{
					item{"Burger", "big size", 55.0},
					item{"Fried Noddle", "instant noddle", 60.0},
				},
			},
		},
		meal{
			Dinner: menu{
				Items: []item{
					item{"Chicken Rice", "chicken & rice", 55.0},
				},
			},
		},
	}

	err := tpl.Execute(os.Stdout, meals)
	if err != nil {
		log.Fatalln(err)
	}
}
