package main

import (
	"log"
	"net/http"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	data := "Electric Electric Bogaloo Bogaloo, that is what my electric bogaloo"
	w.Write([]byte(data))
}

func main() {
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
