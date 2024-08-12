/*
**
Get data from a web server and print out that data
**
*/
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func getDataAndReturnResponse(url string) string {
	// request data from the server using the GET function
	request, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// the data the server returns is in the request body
	defer request.Body.Close()
	// read all the data within the body of the request
	data, err := io.ReadAll(request.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataAndReturnResponse("https://example.com")
	content := []byte(data)
	// log.Println(data)
	// save data to a file
	err := os.WriteFile("response.html", content, 0644)
	if err != nil {
		panic(err)
	}
}
