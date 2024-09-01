/*
**
Exchange greeting messages via channels
**
*/
package main

import "log"

// create the greeter function
func greet(ch chan string) {
	ch <- "Hello"
}

func main() {
	// instantiate the channel and pass it to the greeter
	ch := make(chan string)
	// use a goroutine to pass in the channel
	go greet(ch)
	// pass the return string to the print formatter
	log.Println(<-ch)
}
