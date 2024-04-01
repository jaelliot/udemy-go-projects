package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturn := 5.5

	fmt.Println("Enter investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected return rate: ")
	fmt.Scan(&expectedReturn)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturn/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//fmt.Println(futureValue)
	fmt.Printf("Future Value: %f\nFuture Value (adjusted for Inflation): %f", futureValue, futureRealValue)
	//fmt.Println("Future Value (adjusted for Inflation):", futureRealValue)
}
