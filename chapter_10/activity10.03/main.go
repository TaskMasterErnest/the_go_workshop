/***
Measuring elapsed time
***/

package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	time.Sleep(2 * time.Second)
	end := time.Now()

	elapsedTime := end.Sub(now)

	fmt.Println("The execution took exactly", elapsedTime.Seconds(), "second(s)!")
}
