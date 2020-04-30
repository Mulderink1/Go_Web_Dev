package main

import (
	"github.com/Go_Web_Dev/Basic_Mux/001/middleWare"
	"net/http"
)

func main() {
	http.HandleFunc("/dog/", middleWare.DogRoute)
	http.Handle("/we/", http.HandlerFunc(middleWare.WeRoute))
	http.HandleFunc("/", middleWare.DefaultRoute)

	http.ListenAndServe(":8080", nil)
}
