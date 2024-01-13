package main

import "fmt"

func main() {

	result := add(1, 2)

	fmt.Println(result)
}

// this code basically means - we expect T (type) to be either int, float64 or string
// and we expect a and b to be of type T
// and then the return value will be of type T
// so if we call add(1, 2) the return value will be of type int
// if we call add(1.5, 2.5) the return value will be of type float64
// and if we call add("Hello", "world!") the return value will be of type string
func add[T int | float64 | string](a, b T) T {
	return a + b
}
