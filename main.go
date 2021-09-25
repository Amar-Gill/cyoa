package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type PageData struct {
	PageTitle string
	SomeProp  string
}

func main() {
	tmpl := template.Must(template.ParseFiles("template.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Cool Page Title",
			SomeProp:  "Cool property",
		}
		tmpl.Execute(w, data)
	})
	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", nil)
}
