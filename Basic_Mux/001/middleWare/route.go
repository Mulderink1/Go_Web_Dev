package middleWare

import (
	"html/template"
	"log"
	"net/http"
)

type responseBody struct {
	Name, Message string
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func DefaultRoute(w http.ResponseWriter, r *http.Request) {
	setHeaders(&w)
	tmpData := responseBody{
		"Mike",
		"Default",
	}
	err := tpl.Execute(w, tmpData)
	if err != nil {
		log.Fatalln(err)
	}
}

func WeRoute(w http.ResponseWriter, r *http.Request)  {
	setHeaders(&w)
	tmpData := responseBody{
		"Mike",
		"We",
	}
	err := tpl.Execute(w, tmpData)
	if err != nil {
		log.Fatalln(err)
	}
}

func DogRoute(w http.ResponseWriter, r *http.Request)  {
	setHeaders(&w)
	tmpData := responseBody{
		"Mike",
		"Dog",
	}
	err := tpl.Execute(w, tmpData)
	if err != nil {
		log.Fatalln(err)
	}
}