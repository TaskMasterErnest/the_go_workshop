package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type User struct {
	ID    int64
	Name  string
	Email string
}

func main() {
	users := []User{
		{ID: 43215, Name: "Ernest Joffrey", Email: "kingjoff@lannister.com"},
		{ID: 34687, Name: "Eddard Stark", Email: "manofhonour@winterfell.com"},
	}

	newEmail := "user@packt.com"
	id := int64(34687)
	thisID := int64(43215)

	db, err := sql.Open("postgres", "user=taskmaster password=4@strong-password host=0.0.0.0 port=5432 dbname=postgres sslmode=disable")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The connection to the database has been successfully initialized!")
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Good to go!")
	}

	dbCreate := `
	CREATE TABLE Users
	(
		id integer NOT NULL,
		name text COLLATE pg_catalog."default" NOT NULL,
		email text COLLATE pg_catalog."default" NOT NULL
	)
	WITH (
		OIDS = FALSE
	)
	TABLESPACE pg_default;
	ALTER TABLE Users OWNER to taskmaster;
	ALTER TABLE Users ADD PRIMARY KEY (id);
	`

	// create a table
	_, err = db.Exec(dbCreate)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("The table called Users has been created successfully")
	}

	// insert values into database table
	userCreate := "INSERT INTO Users (id, name, email) VALUES ($1, $2, $3)"
	insertUser, err := db.Prepare(userCreate)
	if err != nil {
		panic(err)
	}
	defer insertUser.Close()
	for _, user := range users {
		_, err := insertUser.Exec(user.ID, user.Name, user.Email)
		if err != nil {
			panic(err)
		} else {
			fmt.Printf("The user with name: %s and email: %s was successfully added.\n", user.Name, user.Email)
		}
	}

	// update/modify some data
	update := "UPDATE Users SET email = $1 WHERE id = $2"
	emailUpdate, err := db.Prepare(update)
	if err != nil {
		panic(err)
	}
	userUpdate, err := emailUpdate.Exec(newEmail, id)
	if err != nil {
		panic(err)
	}
	defer emailUpdate.Close()
	updatedRecords, err := userUpdate.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of records updated:", updatedRecords)
	fmt.Println("The user's email was successfully updated!")

	// delete user from table
	deleteStmt := "DELETE FROM Users WHERE id = $1"
	deleteUser, err := db.Prepare(deleteStmt)
	if err != nil {
		panic(err)
	}
	userDelete, err := deleteUser.Exec(thisID)
	if err != nil {
		panic(err)
	}
	defer deleteUser.Close()
	deletedRecords, err := userDelete.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of records deleted:", deletedRecords)
	fmt.Println("The second user was successfully deleted!")
}
