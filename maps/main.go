package main

import "fmt"

func main() {
	userNames := make([]string, 2, 5)

	userNames = append(userNames, "johndoe")
	userNames = append(userNames, "maxpayne")

	fmt.Println(userNames)

	courseRatings := map[string]float64{}

	courseRatings["go"] = 5.0
	courseRatings["react"] = 4.8

	fmt.Println(courseRatings)
}
