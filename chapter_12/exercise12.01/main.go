/***
In this exercise, we will be catching two signals, SIGINT and SIGTSTP.
Once these signals have been caught, we can then simulate a cleanup of some files.
***/

package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func cleanup() {
	fmt.Println("Simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting files...not really.", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTSTP)
	go func() {
		for {
			s := <- sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted! Someone must have pressed CTRL-C!")
				fmt.Println("Some cleanup is in progress!")
				cleanup()
				done <- true
			case syscall.SIGTSTP:
				fmt.Println()
				fmt.Println("My process has been suspended! Someone pressed CTRL-Z!")
				fmt.Println("Some cleanup is in progress!")
				cleanup()
				done <- true
			}
		}
	}()

	fmt.Println("Program is blocked until a signal is caught: CTRL-C or CTRL-Z")
	<- done
	fmt.Println("Out of here!")
}
