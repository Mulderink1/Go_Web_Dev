package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main (){
	http.HandleFunc("/", dog)
	http.Handle("/images/", image())

	http.ListenAndServe(":8080", nil)
}

func image() http.Handler {
	return http.StripPrefix("/images", http.FileServer(http.Dir("./Images")))
}

func dog(w http.ResponseWriter, r *http.Request){
	method := "nil"

	if r.Method == http.MethodGet {
		method = http.MethodGet
	} else {
		method = "Not a Get request"
	}
	tpl.Execute(w, method)
}