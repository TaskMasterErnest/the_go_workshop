/***
Testing whether a file is being created or not
***/

package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("junk.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("junk.txt: File does not exist!")
			fmt.Println(file)
		}
	}

	fmt.Println()

	file, err = os.Stat("test.txt")
	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println("test.txt: File does not exist!")
			fmt.Println(file)
		}
	}

	fmt.Printf("File Name: %s\nIsDir: %t\nModified: %v\nMode: %v\nSize: %d\n", file.Name(), file.IsDir(), file.ModTime(), file.Mode(), file.Size())

	// read content of file
	content, err := os.ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	// the file content is now a slice of bytes hence for readbility sake, it must be converted to a string.
	fmt.Println("File Contents:")
	fmt.Println(string(content))
}
