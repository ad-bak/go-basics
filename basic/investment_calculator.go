package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Please provide a value: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for Inflation): %.1f`, futureValue, futureRealValue)

}

func calculateFutureValue(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
