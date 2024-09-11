/*
**
We will start a Goroutine with an infinite loop counting from zero until we decide to stop it.
We will make use of the context to notify the routine to stop
We will add a sleep function to make sure we know how many iterations we have executed
**
*/

package main

import (
	"context"
	"log"
	"time"
)

// a function that counts, every 100 ms, from 0
func countNums(c context.Context, n chan int) {
	var v int
	// start loop
	for {
		select {
		// check if context is Done. If yes, pass value to channel and break the loop
		case <-c.Done():
			n <- v
			break
		default:
			// if context is not Done, increment v every 100ms
			time.Sleep(time.Millisecond * 100)
			v++
		}
	}
}

func main() {
	n := make(chan int)
	c := context.TODO()

	cancel, stop := context.WithCancel(c)
	go countNums(cancel, n)

	go func() {
		time.Sleep(time.Millisecond * 100 * 10)
		stop()
	}()

	v := <-n
	log.Println(v)
}
