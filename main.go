package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"

	"eng-cards/templates"

	"github.com/a-h/templ"
)

type Words struct {
	Word    string
	Meaning []string
}

func main() {
	// read file
	bytes, error := os.ReadFile("cards.json")
	if error != nil {
		log.Fatal("error on reading file ", error)
	}

	var data []Words
	error = json.Unmarshal(bytes, &data)
	if error != nil {
		log.Fatal("error to Unmarshall file ", error)
	}

	// get random word
	randomIndex := rand.Intn(len(data))
	log.Printf("random word ", data[randomIndex].Word)

	// create template
	component := templates.Index(data[randomIndex].Meaning, data[randomIndex].Word)

	// run server
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
