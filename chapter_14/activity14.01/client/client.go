package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type Data struct {
	words []string
}

func getDataAndReturnResponse(url string) string {
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()

	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataAndReturnResponse("http://localhost:8000")

	splitWords := strings.Split(data, " ")
	// fmt.Printf("The words split are %v and of type %T\n", splitWords, splitWords)

	wordStruct := Data{
		words: splitWords,
	}

	var countE int
	var countB int

	for _, word := range wordStruct.words {
		word = strings.ToLower(word)
		switch word {
		case "electric":
			countE++
		case "bogaloo":
			countB++
		}
	}
	fmt.Println("Electric Count: ", countE)
	fmt.Println("Bogaloo Count: ", countB)
}
