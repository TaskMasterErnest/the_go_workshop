/*
**
The client to read data from the server
**
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type messageData struct {
	Message string `json:"message"`
}

func getDataAndReturnResponse(url string) messageData {
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer request.Body.Close()

	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	// parse the response into JSON
	// initialize an empty struct od messageData
	message := messageData{}
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Fatal(err)
	}

	return message
}

func main() {
	data := getDataAndReturnResponse("http://localhost:8000")
	fmt.Println(data.Message)
}
