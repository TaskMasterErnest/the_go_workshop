package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

func readFiles(filename string, out chan int, wg *sync.WaitGroup) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	// set a buffer reader
	reader := bufio.NewReader(file)

	// loop through the byte and read by line
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				wg.Done()
				return
			} else {
				panic(err)
			}
		}
		s := strings.ReplaceAll(str, "\n", "")
		// since they are numbers, convert them to int
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		// pass the int to the out channel
		out <- i
	}
}

func sortInput(in chan int, odd chan int, even chan int, wg *sync.WaitGroup) {
	for i := range in {
		switch i % 2 {
		case 0:
			even <- i
		case 1:
			odd <- i
		}
	}
	close(even)
	close(odd)

	wg.Done()
}

func calcInput(in chan int, out chan int, wg *sync.WaitGroup) {
	var sum int
	for i := range in {
		sum += i
	}
	// pass the sum to the out channel
	out <- sum
	wg.Done()
}

func writeToFile(even, odd chan int, wg *sync.WaitGroup, resultFile string) {
	newFile, err := os.Create(resultFile)
	if err != nil {
		panic(err)
	}

	// loop through the incoming input
	for i := 0; i < 2; i++ {
		select {
		case i := <-even:
			newFile.Write([]byte(fmt.Sprintf("Even: %d\n", i)))
		case i := <-odd:
			newFile.Write([]byte(fmt.Sprintf("Odd: %d\n", i)))
		}
	}
	wg.Done()
}

func main() {
	// create waitgroups
	wg := &sync.WaitGroup{}
	wg.Add(2)

	wg2 := &sync.WaitGroup{}
	wg2.Add(4)

	// initialize the channels
	odd := make(chan int)
	even := make(chan int)
	out := make(chan int)
	sumOdd := make(chan int)
	sumEven := make(chan int)

	// run the routines
	go readFiles("./file1.txt", out, wg)
	go readFiles("./file2.txt", out, wg)

	go sortInput(out, odd, even, wg2)

	go calcInput(even, sumEven, wg2)
	go calcInput(odd, sumOdd, wg2)

	go writeToFile(sumEven, sumOdd, wg2, "./result.txt")

	wg.Wait()
	close(out)
	wg2.Wait()
}
