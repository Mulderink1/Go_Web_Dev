package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type JSON struct {
	Works       int `json:"Works"`
	StatusFound int `json:"StatusFound"`
	NotFound    int `json:"NotFound"`
}

func main() {
	var data JSON

	rcvd := `{"Works": 200, "StatusFound": 303, "NotFound": 404}`

	err := json.Unmarshal([]byte(rcvd), &data)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(data)
	fmt.Println(data.StatusFound)
}
