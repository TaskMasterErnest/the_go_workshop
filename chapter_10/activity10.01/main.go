/***
A script to format a date according to user requirements
***/

package main

import (
	"fmt"
	"strconv"
	"time"
)

var (
	date          = time.Now()
	formattedDate = strconv.Itoa(date.Hour()) + ":" + strconv.Itoa(date.Minute()) + ":" + strconv.Itoa(date.Second()) + "  " + strconv.Itoa(date.Year()) + "/" + strconv.Itoa(int(date.Month())) + "/" + strconv.Itoa(date.Day())
)

func main() {
	fmt.Println("The time is:", formattedDate)
}
