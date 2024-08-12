package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	filename = "data.json"
)

type nameData struct {
	Names []string `json:"name"`
}

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		decoder := json.NewDecoder(r.Body)
		nameData := nameData{}
		err := decoder.Decode(&nameData)
		if err != nil {
			log.Fatal(err)
		}
		dataBytes, err := json.Marshal(nameData)
		if err != nil {
			log.Fatal(err)
		}

		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		err = os.WriteFile(filename, dataBytes, 0644)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodGet:
		file, err := os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		w.Write(data)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}

}

func main() {
	fmt.Println("Starting server on port :8000")
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
