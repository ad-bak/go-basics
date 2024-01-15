package main

import "fmt"

func main() {
	numbers := []int{1, 10, 15}

	//2. this is how we pass a slice to a variadic function
	// we use the ... operator after the slice name
	// or we could simply pass the numbers slice like this:
	// sum := sumup(numbers)
	sum := sumup(10, numbers...)

	fmt.Println(sum)
}

// 1. this type of syntax is called variadic function
// it means that we can pass any number of arguments to this function
func sumup(startingValue int, numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
