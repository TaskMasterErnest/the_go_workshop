/***
In this function, we will create a function called WhatsTheClock.
The goal of this is to create a function that wraps a nice, formatted time.Now() function and returns
the date in ANSIC format
The ANSIC format will be explained in further detail
***/

package main

import (
	"fmt"
	"time"
)

func WhatsTheClock() string {
	return time.Now().Format(time.ANSIC)
}

func main() {
	fmt.Println(WhatsTheClock())
}
