package main

import (
	"fmt"
	"math"
)

func main() {

	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
	var inflationRate float64

	print("investment amount: ")
	fmt.Scan(&investmentAmount)

	print("years: ")
	fmt.Scan(&years)

	print("expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	print("inflation rate: ")
	fmt.Scan(&inflationRate)

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

	fmt.Printf("investment amount %.2f\n", investmentAmount)
	fmt.Printf("future value: %.2f\n", futureValue)
	fmt.Printf("future real value %.2f\n", futureRealValue)
}
