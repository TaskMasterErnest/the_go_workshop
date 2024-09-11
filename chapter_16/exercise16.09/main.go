/*
**
In this exercise, we will have one routine to send messages and another one to print them.
We want to know when the sender has finished sending messages
**
*/

package main

import "log"

// define function that receive strings and print them later
func readThem(in, out chan string) {
	// loop over the contents of the in channel
	for i := range in {
		log.Println(i)
	}
	// send notification when done
	out <- "done"
}

func main() {
	// set log flags to 0 so we do not see anything other than the strings.
	log.SetFlags(0)

	// instantiate the necessary channels
	in := make(chan string)
	out := make(chan string)

	// use goroutine to spin the channels up
	go readThem(in, out)

	sts := []string{"a", " b", "  c", "   d", "    e", "     f"}
	for _, v := range sts {
		// send them to the in channel
		in <- v
	}
	close(in)

	// output the content of the out channel
	<-out
}
