package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

func postFileAndReturnResponse(filename string, url string) string {
	// create buffer to write file bytes into
	fileDataBuffer := bytes.Buffer{}
	// a writer to allow bytes to write into buffer
	multipartWriter := multipart.NewWriter(&fileDataBuffer)

	// open file on local computer
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	// create a formfile
	// wrap data in the right format to upload to the server
	formFile, err := multipartWriter.CreateFormFile("myFile", file.Name())
	if err != nil {
		log.Fatal(err)
	}

	// copy bytes from local file (on PC) to formFile and close it
	_, err = io.Copy(formFile, file)
	if err != nil {
		log.Fatal(err)
	}
	multipartWriter.Close()

	// now we need to create a POST request.
	// we need more control over the data being sent hence we use an http.Request instead of an http.Post
	// the bytes buffer we created also need to be in the request
	request, err := http.NewRequest("POST", url, &fileDataBuffer)
	if err != nil {
		log.Fatal(err)
	}

	// set the Content-Type header so that the server knows how to handle the request
	request.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// send the request as follows
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// read data from the response
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}

func main() {
	var filePath string
	fmt.Print("Path to file to upload: ")
	fmt.Scan(&filePath)
	data := postFileAndReturnResponse(filePath, "http://localhost:8000")
	fmt.Println(data)
}
