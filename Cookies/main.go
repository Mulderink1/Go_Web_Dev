package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("counter")
		if err == http.ErrNoCookie {
			cookie = &http.Cookie{
				Name:  "counter",
				Value: "0",
			}
		}

		count, err := strconv.Atoi(cookie.Value)
		if err != nil {
			log.Fatal(err)
		}

		count++
		cookie.Value = strconv.Itoa(count)
		http.SetCookie(w, cookie)
		fmt.Fprint(w, fmt.Sprintf(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Cookie Value: %d</strong></body></html>`, count))
	})

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFoundHandler()
	})

	http.ListenAndServe(":8080", nil)
}
