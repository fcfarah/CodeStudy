package main

import "fmt"

func appV1() {
	var revenue float64
	var expenses float64
	var taxRate float64

	print("revenue: ")
	fmt.Scan(&revenue)
	print("expenses: ")
	fmt.Scan(&expenses)
	print("tax rate: ")
	fmt.Scan(&taxRate)
	taxRate = taxRate / 100

	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate)
	profit := earningsAfterTax

	ratio := earningsBeforeTax / profit

	fmt.Printf("earnings before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf(`ratio: %.2f
	
	\n`, ratio)

}
