/***
A basic server that recieves data in multipart form and saves the file within the form
***/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type server struct{}

func (srv server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	uploadedFile, uploadedFileHeader, err := r.FormFile("myFile")
	if err != nil {
		log.Fatal(err)
	}
	defer uploadedFile.Close()

	fileContent, err := io.ReadAll(uploadedFile)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(fmt.Sprintf("./%s", uploadedFileHeader.Filename), fileContent, 0666)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("%s Uploaded!", uploadedFileHeader.Filename)))
}

func main() {
	fmt.Println("Starting server on port :8000...")
	log.Fatal(http.ListenAndServe(":8000", server{}))
}
