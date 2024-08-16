package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type PageWithCounter struct {
	Heading string `json:"heading"`
	Content string `json:"content"`
	Counter int    `json:"views"`
}

func (h *PageWithCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Counter++
	encode, err := json.MarshalIndent(h, "", "    ")
	if err != nil {
		fmt.Println("error marshalling json: ", err)
	}
	w.Write(encode)
}

func main() {
	main := PageWithCounter{Heading: "Hello World!", Content: "This is the main page."}
	cha1 := PageWithCounter{Heading: "Chapter 1", Content: "This is the first chapter."}
	cha2 := PageWithCounter{Heading: "Chapter 2", Content: "This is the second chapter"}

	http.Handle("/", &main)
	http.Handle("/chapter1", &cha1)
	http.Handle("/chapter2", &cha2)
	fmt.Println("Starting server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
