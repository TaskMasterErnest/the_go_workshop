/***
Show the number of prime numbers in the database
Show the numbers in order of appearance
Sum the prime numbers
Remove all even numbers from the table
***/

package main

import (
	"database/sql"
	"fmt"
	"math/big"

	_ "github.com/lib/pq"
)

func main() {
	var number int64
	var prop string
	var primeSum int64
	var newNumber int64

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

	// get a list of all the prime numbers
	allTheNumbers := "SELECT * FROM Number"
	Numbers, err := db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	primeSum = 0
	result, err := Numbers.Query()
	if err != nil {
		panic(err)
	}
	fmt.Println("The list of prime numbers:")
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		if big.NewInt(number).ProbablyPrime(0) {
			primeSum += number
			fmt.Println(" ", number)
		}
	}
	Numbers.Close()
	fmt.Println("\nThe total sum of prime numbers in this range is:", primeSum)

	// remove the even numbers
	remove := "DELETE FROM Number WHERE Property=$1"
	removeResult, err := db.Exec(remove, "Even")
	if err != nil {
		panic(err)
	}
	removedRecords, err := removeResult.RowsAffected()
	if err != nil {
		panic(err)
	}
	fmt.Println("The number of rows removed:", removedRecords)

	// updating numbers
	fmt.Println("Updating numbers...")
	update := "UPDATE Number SET Number=$1 WHERE Number=$2 and Property=$3"
	allTheNumbers = "SELECT * FROM Number"
	Numbers, err = db.Prepare(allTheNumbers)
	if err != nil {
		panic(err)
	}
	result, err = Numbers.Query()
	if err != nil {
		panic(err)
	}
	for result.Next() {
		err = result.Scan(&number, &prop)
		if err != nil {
			panic(err)
		}
		newNumber = number + primeSum
		_, err = db.Exec(update, newNumber, number, prop)
		if err != nil {
			panic(err)
		}
	}
	Numbers.Close()
	fmt.Println("The execution is now complete...")

	db.Close()
}
