package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Amar-Gill/cyoa"
)

func main() {
	f, err := os.Open("gopher.json")
	if err != nil {
		panic(err)
	}

	var story cyoa.Story

	d := json.NewDecoder(f)

	d.Decode(&story)

	h := cyoa.NewHandler(story)

	fmt.Println("Listening on port :8080")
	http.ListenAndServe(":8080", h)
}
