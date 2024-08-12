/*
**
This is the server.
We will serve some data from here that we will want the client to access
**
*/
package main

import (
	"log"
	"net/http"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	message := "{\"message\": \"hello world\"}"
	w.Write([]byte(message))
}

func main() {
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
