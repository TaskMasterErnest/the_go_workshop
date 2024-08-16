/***
Creating a Hello World server using static files.
***/

package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Starting server on port :8000...")
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "./exercise15.05/index.html")
	// })
	// ditching ServeFile for FileServer, for more performance.
	fileServer := http.FileServer(http.Dir("./exercise15.05"))
	http.Handle("/", fileServer)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
