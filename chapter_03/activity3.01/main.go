package main

import "fmt"

type SalesTax struct {
	Item    string
	Cost    float64
	TaxRate float64
}

func Calculate(item *SalesTax) float64 {
	// calculae the sales tax rate for the item.
	// sales tax = cost * sales tax rate
	totalTax := item.Cost * (item.TaxRate / 100)
	return totalTax
}

func main() {
	// item1 := SalesTax{
	// 	Item:    "Cake",
	// 	Cost:    0.99,
	// 	TaxRate: 7.5,
	// }

	// // calulate sales tax rate
	// rate := Calculate(&item1)
	// fmt.Printf("You pay %.2f tax for %s\n", rate, item1.Item)

	items := []SalesTax{
		{
			Item:    "Cake",
			Cost:    0.99,
			TaxRate: 7.5,
		},
		{
			Item:    "Milk",
			Cost:    2.75,
			TaxRate: 1.5,
		},
		{
			Item:    "Butter",
			Cost:    0.87,
			TaxRate: 2,
		},
	}

	var TotalTaxPayable float64
	for _, item := range items {
		rate := Calculate(&item)
		TotalTaxPayable += rate
		fmt.Printf("You pay %.2f tax for %s\n", rate, item.Item)
	}
	fmt.Printf("You pay a total tax of %.2f on the goods purchased\n", TotalTaxPayable)
}
