/***
basic web server that receives a JSON post request and returns the message sent to it back to the client
***/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct{}

type messageData struct {
	Message string `json:"message"`
}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// decode the message body from client
	jsonDecoder := json.NewDecoder(r.Body)
	messageData := messageData{}
	err := jsonDecoder.Decode(&messageData)
	if err != nil {
		log.Fatal(err)
	}
	jsonBytes, err := json.Marshal(messageData)
	log.Println(string(jsonBytes))
	w.Write(jsonBytes)
}

func main() {
	fmt.Println("Starting server on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
