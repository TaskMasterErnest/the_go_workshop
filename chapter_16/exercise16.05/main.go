/***
Two-Way message exchange with channels
***/

package main

import (
	"fmt"
	"log"
)

// write the greeter function
// this will first wait for a message, and then reply with a thank-you and then send its own greeting
func greet(ch chan string) {
	msg := <-ch
	ch <- fmt.Sprintf("Thanks for the message: \"%s\"", msg)
	ch <- "Hello David"
}

func main() {
	ch := make(chan string)
	go greet(ch)

	// send message
	ch <- "Hello John"

	// log the messages to get back (we will use a loop later)
	log.Println(<-ch)
	log.Println(<-ch)
}
