/***
Practicing with the os.OpenFile command
***/

package main

import "os"

func main() {
	file, err := os.OpenFile("junk101.txt", os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file2, err := os.OpenFile("junk102.txt", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file2.Close()
	// writing to the file
	if _, err := file2.Write([]byte("Adding junk stuff.\n")); err != nil {
		panic(err)
	}

	file3, err := os.OpenFile("junk102.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file3.Close()
	// appending to file
	if _, err := file3.Write([]byte("Adding MORE junk stuff.\n")); err != nil {
		panic(err)
	}
}
