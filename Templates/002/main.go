package main

import (
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Region string
	Hotel []hotel
}

type regions []region

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("hw.gohtml"))
}

func main() {
	caHotels := regions{
		region{
			"Central",
			[]hotel{
				{
					Name:    "Mike",
					Address: "LA",
					City:    "Venice",
					Zip:     "99999",
				},
			},
		},
	}
	err := tpl.Execute(os.Stdout, caHotels)
	if err != nil {
		log.Fatalln(err)
	}
}