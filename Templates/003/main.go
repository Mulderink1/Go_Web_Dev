package main

import (
	"log"
	"os"
	"text/template"
)

type breakfast struct {
	Name string
}

type lunch struct {
	Name string
}
type dinner struct {
	Name string
}

type restaurant struct {
	Breakfast []breakfast
	Lunch []lunch
	Dinner []dinner
}

type restaurants []restaurant

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("hw.gohtml"))
}

func main() {
	restaurantSlice := restaurants{
		restaurant{
			[]breakfast{
				{
					"PanCakes",
				},
			},
			[]lunch{
				{
					"HotDog",
				},
			},
			[]dinner{
				{
					"Polish",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, restaurantSlice)
	if err != nil {
		log.Fatalln(err)
	}
}