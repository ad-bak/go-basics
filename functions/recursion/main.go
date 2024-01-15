package recursion

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

}

// RECURSIVE SOLUTION
func factorial(number int) int {
	if number == 1 {
		return 1
	}

	return number * factorial(number-1)
}

// LOOP BASED SOLUTION
// func factorial(number int) int {
// 	result := 1

// 	for i := 1; i <= number; i++ {
// 		result *= i
// 	}

// 	return result
// }
