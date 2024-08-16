/***
A simple hello server that returns hello world when queried
***/

package main

import (
	"log"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World</h1>"
	w.Write([]byte(msg))
}

func main() {
	log.Fatal(http.ListenAndServe(":8000", Hello{}))
}
