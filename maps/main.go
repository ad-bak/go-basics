package main

import "fmt"

// type alias that will be used in the main function
// we are saying that floatMap is a map that has string as key and float64 as value
type floatMap map[string]float64

// here we are adding a method to the floatMap type
func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	//here we are allocating memory for 2 strings and we are saying that the capacity of the array is 5
	// if we overpass the capacity of the array, go will allocate more memory for us
	// but if we know in advance how many elements we will have in the array, we can allocate memory for them
	// and this will be more efficient
	userNames := make([]string, 2, 5)

	userNames[0] = "tonymontana"

	userNames = append(userNames, "johndoe")
	userNames = append(userNames, "maxpayne")

	fmt.Println(userNames)

	// same as above, but we are allocating memory for 3 strings and we are saying that the capacity of the array is 3
	// when we append the third element, go will allocate more memory for us
	courseRatings := make(floatMap, 3)

	courseRatings["go"] = 5.0
	courseRatings["react"] = 4.8
	courseRatings["python"] = 4.7
	courseRatings["ruby"] = 4.5

	courseRatings.output()

	for index, value := range userNames {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

	for key, value := range courseRatings {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}
}
