package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// a list of Hellos from other languages
var helloList = []string{
	"Hello, world",
	"Καλημέρα κόσμε",
	"こんにちは世界",
	"سلام دنیا‎",
	"Привет, мир",
}

func main() {
	// seed a random number generator using the current time
	rand.NewSource(time.Now().UnixNano())
	// generate a random number in the range of the list
	index := rand.Intn(len(helloList))
	// call a function and receive multiple return values
	msg, err := hello(index)
	// handle any errors
	if err != nil {
		log.Fatal(err)
	}
	// print message to the console
	fmt.Println(msg)
}

// consider the hello() function
func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		// create an error and convert the int type to a string
		return "", errors.New("out of range: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
