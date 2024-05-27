package main

import (
	"fmt"
	"os"
	"strings"
)

// define a locale struct with a lnaguage and a territory
type locale struct {
	language  string
	territory string
}

// function that returns the test data
func getLocales() map[locale]struct{} {
	supportedLocales := make(map[locale]struct{}, 5)
	supportedLocales[locale{"en", "US"}] = struct{}{}
	supportedLocales[locale{"en", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "CN"}] = struct{}{}
	supportedLocales[locale{"fr", "FR"}] = struct{}{}
	supportedLocales[locale{"ru", "RU"}] = struct{}{}

	return supportedLocales
}

// function to use a passed locale struct to check against the sample data, if a locale exists
func localeExists(l locale) bool {
	_, exists := getLocales()[l]
	return exists
}

func main() {
	// check argument length passed
	if len(os.Args) < 2 {
		fmt.Println("No locale passed")
		os.Exit(1)
	}

	// process the passed locale struct to confirm it is in a valid format
	localeParts := strings.Split(os.Args[1], "_")
	if len(localeParts) != 2 {
		fmt.Printf("Invalid locale passed: %v\n", os.Args[1])
		os.Exit(1)
	}

	// create a locale struct with the passed arguments
	passedLocale := locale{
		territory: localeParts[1],
		language:  localeParts[0],
	}

	if !localeExists(passedLocale) {
		fmt.Printf("Locale not supported: %v\n", os.Args[1])
		os.Exit(1)
	}

	fmt.Println("Locale passed is supported")
}
