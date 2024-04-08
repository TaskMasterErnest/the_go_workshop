package main

import (
	"fmt"
)

func main() {
	// define a map and initialize it
	config := map[string]string{
		"debug":    "1",
		"loglevel": "warn",
		"version":  "1.2.2",
	}

	// loop over them with a range
	for key, value := range config {
		fmt.Println(key, "=", value)
	}
}
