/***
Create an HTTP client and set custom options on it.
***/

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	auth := r.Header.Get("Authorization")
	if auth != "superSecretToken" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Authorization token not recognized!"))
		return
	}
	time.Sleep(5 * time.Second)
	message := "Hello Client!"
	w.Write([]byte(message))
}

func main() {
	fmt.Println("Starting server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
