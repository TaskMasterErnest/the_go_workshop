/***
Create code to backup files from an original file to a backup file
existing file is note.txt, backup is backupFile.txt.
***/

package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

var (
	ErrExistingFileNotFound = errors.New("The existing file is not found")
)

func createBackup(existing, backup string) error {
	// check to see if existing file exists before backing up
	_, err := os.Stat(existing)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrExistingFileNotFound
		}
		return err
	}

	workingFile, err := os.Open(existing)
	if err != nil {
		return err
	}
	// check to see whether there is an error
	content, err := io.ReadAll(workingFile)
	if err != nil {
		return err
	}

	// write to the backup file
	err = os.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}

func addNotes(workingFile, notes string) error {
	notes += "\n"
	file, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// write contents of notes to working file
	if _, err := file.Write([]byte(notes)); err != nil {
		return err
	}

	return nil
}

func main() {
	backupFile := "backupFile.txt"
	workingFile := "note.txt"
	data := "note"

	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for i := 0; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile, note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}
