package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func getDataWithCustomOptionsAndReturnResponse(url string, authToken string) string {
	// set timeout to +1 second more
	client := http.Client{Timeout: 6 * time.Second}

	// create a GET request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	// account for the request headers. Set the authorization header
	request.Header.Set("Authorization", authToken)
	// have the default http client do the request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	data, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func main() {
	data := getDataWithCustomOptionsAndReturnResponse("http://localhost:8000", "superSecretToken")
	fmt.Println(data)
}
