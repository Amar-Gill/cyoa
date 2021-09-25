package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Chapter struct {
	Title   string
	Story   []string
	Options []Option
}

type Option struct {
	Text    string
	Chapter string
}
type PageData struct {
	Intro     Chapter
	NewYork   Chapter
	Debate    Chapter
	SeanKelly Chapter
	MarkBates Chapter
	Home      Chapter
}

func main() {
	jsonBytes, err := os.ReadFile("gopher.json")
	if err != nil {
		panic(err)
	}

	var data PageData

	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("template.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, data)
	})
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
