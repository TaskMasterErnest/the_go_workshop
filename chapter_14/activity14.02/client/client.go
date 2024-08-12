package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type postName struct {
	Names []string `json:"name"`
}

type getName struct {
	Names []string `json:"name"`
}

func addNameAndParseResponse(url string, names postName) {
	nameBytes, err := json.Marshal(names)
	if err != nil {
		log.Fatal(err)
	}

	request, err := http.Post(url, "application/json", bytes.NewBuffer(nameBytes))
	if err != nil {
		log.Fatal(err)
	}

	defer request.Body.Close()
}

func getDataAndParseResponse(url string) ([]string, error) {
	request, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer request.Body.Close()

	data, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}

	var response getName
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return response.Names, nil
}

func main() {
	putName := postName{Names: []string{"Ernest", "Clinton", "Gloria", "Bernice"}}
	addNameAndParseResponse("http://localhost:8000", putName)
	fmt.Println("<<<------------------------------------------------------>>>")
	data, err := getDataAndParseResponse("http://localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for _, name := range data {
		fmt.Println(name)
	}
}
