/***
We will build a server that accepts a JSON message and will respond with another JSON message.
We will use Postman to test the functionality
***/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Request struct {
	Name    string
	Surname string
}

type Response struct {
	Greeting string
}

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var data Request
	err := decoder.Decode(&data)
	if err != nil {
		w.WriteHeader(400)
		return
	}

	response := Response{Greeting: fmt.Sprintf("Greetings, %s %s.", data.Name, data.Surname)}
	bytes, err := json.MarshalIndent(response, "", "    ")
	if err != nil {
		w.WriteHeader(400)
		return
	}
	w.Write(bytes)
}

func main() {
	http.Handle("/", Hello{})
	fmt.Println("Starting server at port :8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
