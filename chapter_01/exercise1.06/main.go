package main

import (
	"fmt"
	"time"
)

// create a function to get a configuration
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func main() {
	// assign variables to the returned values
	// the varibles capture the returned values
	Debug, LogLevel, startUpTime := getConfig()

	// print them out
	fmt.Println(Debug, LogLevel, startUpTime)
}
