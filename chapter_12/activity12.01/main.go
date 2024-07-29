/***
Ingest a transaction file from a bank
File is a csv file.
Create a CLI that will accept two flags
***/

package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	ErrBudgetCategoryNotFound = errors.New("budget category not found")
)

type budgetCategory string

const (
	autoFuel   budgetCategory = "fuel"
	food       budgetCategory = "food"
	mortgage   budgetCategory = "mortgage"
	repairs    budgetCategory = "repairs"
	insurance  budgetCategory = "insurance"
	utilities  budgetCategory = "utilities"
	retirement budgetCategory = "retirement"
)

type Transaction struct {
	ID       string
	Payee    string
	Spent    float64
	Category budgetCategory
}

func matchBudgetCategories(value string) (budgetCategory, error) {
	// assert whether the string is a certain one, map that string to budget category
	value = strings.TrimSpace(strings.ToLower(value))
	switch value {
	case "fuel", "gas":
		return autoFuel, nil
	case "car insurance", "life insurance":
		return insurance, nil
	case "mortgage":
		return mortgage, nil
	case "food":
		return food, nil
	case "repairs":
		return repairs, nil
	case "utilities":
		return utilities, nil
	case "retirement":
		return retirement, nil
	default:
		return "", ErrBudgetCategoryNotFound
	}
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("File: %s does not exist\n", filename)
			os.Exit(1)
		}
	}
	return nil
}

func parseTxCSVFile(bankTransactions io.Reader, logFile string) []Transaction {
	r := csv.NewReader(bankTransactions)
	transactions := []Transaction{}
	// manage the headers
	header := true

	for {
		transaction := Transaction{}
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		// continue managing headers
		if !header {
			for idx, value := range record {
				switch idx {
				case 0:
					// id
					value = strings.TrimSpace(value)
					transaction.ID = value
				case 1:
					// payee
					value = strings.TrimSpace(value)
					transaction.Payee = value
				case 2:
					// spent
					value = strings.TrimSpace(value)
					transaction.Spent, err = strconv.ParseFloat(value, 64)
					if err != nil {
						log.Fatal(err)
					}
				case 3:
					//category
					transaction.Category, err = matchBudgetCategories(value)
					if err != nil {
						s := strings.Join(record, ", ")
						writeErrorToLog("error converting CSV category column - ", err, s, logFile)
						fmt.Println(err)
					}
				}
			}
			transactions = append(transactions, transaction)
		}
		header = false
	}
	return transactions
}

func writeErrorToLog(msg string, err error, data string, logFile string) error {
	msg += "\n" + err.Error() + "\nData: " + data + "\n\n"
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// write to file
	if _, err := file.Write([]byte(msg)); err != nil {
		return err
	}
	return nil
}

func main() {
	txFile := flag.String("csv", "", "path to CSV file")
	logFile := flag.String("logfile", "", "path to log file")
	flag.Parse()

	//require the CSV file
	if *txFile == "" {
		fmt.Println("csv file is required")
		flag.PrintDefaults()
		os.Exit(1)
	}
	// require the log file
	if *logFile == "" {
		fmt.Println("logfile location is required")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// check if stated files exist
	checkFile(*txFile)
	checkFile(*logFile)

	// work on CSV file
	file, err := os.Open(*txFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	txs := parseTxCSVFile(file, *logFile)
	fmt.Println()
	for _, tx := range txs {
		fmt.Printf("%v\n", tx)
	}

}
