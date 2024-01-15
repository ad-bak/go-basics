package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}

	// here we use a closure
	// we create a function that will multiply a number by 2
	double := createTransformer(2)
	triple := createTransformer(3)

	transformed := transformNumbers(&numbers, func(number int) int {
		return number * 2
	})

	// then we pass the function to the transformNumbers function
	// since we already have a function that will multiply a number by 2
	// we can just pass it to the transformNumbers function
	doubled := transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)

	fmt.Println("double", doubled)
	fmt.Println("triple", tripled)

	fmt.Println(transformed)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}

	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}

	return dNumbers
}

// this is a closure
// it's a function that returns another function
// in this example this function will return a function that will multiply a number by the factor
func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
