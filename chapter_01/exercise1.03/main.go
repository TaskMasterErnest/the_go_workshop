package main

import (
	"fmt"
	"time"
)

// assign variables
var (
	Debug       bool      = false
	LogLevel    string    = "info"
	startUpTime time.Time = time.Now()
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
