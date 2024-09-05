package main

import (
	"bytes"
	"fmt"
	"os"
)

func readInput(in chan []byte, out chan []byte) {
	var buffer bytes.Buffer
	var data []byte
	for d := range in {
		for _, byteValue := range d {
			err := buffer.WriteByte(byteValue)
			if err != nil {
				fmt.Println("error writing to buffer:", err)
				return
			}
		}
		data = buffer.Bytes()
		out <- data
	}
}

func readFiles(files ...string) error {
	for _, file := range files {
		data, err := os.ReadFile(file)
		if err != nil {
			panic(err)
		}

		// invoke the writeToFile
		writeToFile(data)
	}
	return nil
}

func writeToFile(data []byte) error {
	// write the data of the files to new file
	newFile, err := os.OpenFile("result.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer newFile.Close()
	if _, err := newFile.Write(data); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	readFiles("./file1.txt", "./file2.txt")
}
