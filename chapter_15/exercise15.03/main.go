/***
In this exercise, we will again create an HTTP server that is able to cheer us,
but instead of a general "hello world" message, we will provide a message depending on our name.
The idea is that, by opening the browser on the server's URL and adding a parameter called name,
the server will welcome us with the message "hello ", followed by the value of the name parameter.
The server is very simple and does not have sub-pages,
but contains this dynamic element that constitutes a starting point for more complex situations
***/

package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	val := r.URL.Query() // the Query() returns a map[string][]string for all the parameters in the querystring
	// get the value of a specific parameter called name
	name, ok := val["name"]
	// check if the name key exists in the parameter
	// if name key does not exist, it prompts a 400 status code and a response message
	// the response message states that the name parameter has not been sent.
	if !ok {
		w.WriteHeader(400)
		w.Write([]byte("Missing name"))
		return
	}

	w.Write([]byte(fmt.Sprintf("Hello %s", strings.Join(name, ","))))
}

func main() {
	http.HandleFunc("/", Hello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
