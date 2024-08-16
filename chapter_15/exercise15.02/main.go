/***
We will introduce a couple of routes we acn use
**/

package main

import (
	"log"
	"net/http"
)

type Hello struct{}
type Chapter1 struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Hello World!</h1>"
	w.Write([]byte(msg))
}

func (c1 Chapter1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	msg := "<h1>Chapter One(1)</h1>"
	w.Write([]byte(msg))
}

func main() {
	http.Handle("/", Hello{})
	http.Handle("/chapter1", Chapter1{})
	log.Fatal(http.ListenAndServe(":8000", nil))
}
