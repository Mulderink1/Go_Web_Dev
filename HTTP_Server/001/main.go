package main

import (
	"html/template"
	"io"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("dog.gohtml"))
}

func main (){
	http.HandleFunc("/", foo)
	http.HandleFunc("/dog/", dog)
	http.Handle("/images/", image())

	http.ListenAndServe(":8080", nil)
}

func image() http.Handler {
	return http.StripPrefix("/images", http.FileServer(http.Dir("./Images")))
}

func foo(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, `<h1>Foo Ran</h1>`)
}

func dog(w http.ResponseWriter, r *http.Request){
	tpl.Execute(w, nil)
}