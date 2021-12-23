package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	h := cyoa.NewHandler(story)

	fmt.Println("Listening on port :8080")
	log.Fatal(http.ListenAndServe(":8080", h))
}
