package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	var prop string

	db, err := sql.Open("postgres", "user=taskmaster password=4@strong-password host=0.0.0.0 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the database was successfully initialized!")
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Good to go!")
	}

	TableCreate := `
	CREATE TABLE Number
	(
		Number integer NOT NULL,
		Property text COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Number OWNER to taskmaster;
	`

	_, err = db.Exec(TableCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table has been created created!")
	}

	insert, insertErr := db.Prepare("INSERT INTO Number VALUES ($1, $2)")
	if insertErr != nil {
		panic(err)
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			prop = "Even"
		} else {
			prop = "Odd"
		}
		_, err = insert.Exec(i, prop)
		if err != nil {
			panic(err)
		} else {
			fmt.Println("The number:", i, "is:", prop)
		}
	}

	insert.Close()
	fmt.Println("The numbers are ready!")

	db.Close()

}
