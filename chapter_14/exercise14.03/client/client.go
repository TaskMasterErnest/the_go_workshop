package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func postDataAndReturnResponse(msg messageData, url string) messageData {
	jsonBytes, _ := json.Marshal(msg)

	request, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()

	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	msg := messageData{Message: "Hello there! Server"}
	data := postDataAndReturnResponse(msg, "http://localhost:8000")
	fmt.Println(data.Message)
}
