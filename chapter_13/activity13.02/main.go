package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type User struct {
	UserID  int64
	Message []string
}

func main() {
	messages := []User{
		{UserID: 42357, Message: []string{"Off with his head", "The Hand of the King's head", "I have never said such words"}},
		{UserID: 52551, Message: []string{"Even superheroes die", "Bad lettuce"}},
		{UserID: 90083, Message: []string{"Oh my Lord", "Praise his name!"}},
	}

	db, err := sql.Open("postgres", "user=taskmaster password=4@strong-password host=0.0.0.0 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the database has been successfully initialized!")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Good to go!")
	}

	defer db.Close()

	tableCreate := `
	CREATE TABLE Messages
	(
		userId INT NOT NULL,
		message VARCHAR(280)
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Messages OWNER to taskmaster;
	`

	// create the messages table
	_, err = db.Exec(tableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table Messages has been successfully created!")
	}
	fmt.Println("--------------------------------------------")
	fmt.Println()

	// load messages to Messages table
	loadStmt := "INSERT INTO Messages (userId, message) VALUES ($1, $2)"
	prepLoad, err := db.Prepare(loadStmt)
	if err != nil {
		panic(err)
	}
	defer prepLoad.Close()
	for _, msg := range messages {
		for i, message := range msg.Message {
			_, err := prepLoad.Exec(msg.UserID, message)
			if err != nil {
				panic(err)
			} else {
				fmt.Printf("The user with ID: %d with message: '%s' has been successfully added.\n", msg.UserID, msg.Message[i])
			}
		}
	}

	// query the Messages
	queryStmt := "SELECT message FROM Messages WHERE userId = $1"
	loadQuery, err := db.Prepare(queryStmt)
	if err != nil {
		panic(err)
	}
	defer loadQuery.Close()

	var message string
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("--------------------------------------------")
	fmt.Println()
	fmt.Print("Give me the userId: ")
	uidStr, err := reader.ReadString('\n')
	uidStr = strings.TrimRight(uidStr, "\r\n")
	if err != nil {
		panic(err)
	}
	uid, err := strconv.ParseInt(uidStr, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println("Looking for messages for user with ID:", uid, "##")
	result, err := loadQuery.Query(uid)
	if err != nil {
		fmt.Printf("Query returned nothing. No such user with ID: %v.\n", err)
	}
	for result.Next() {
		err = result.Scan(&message)
		if err != nil {
			panic(err)
		}
		fmt.Println("The user with ID:", uid, "sent the following message:", message)
	}

}
