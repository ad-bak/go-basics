package main

import "fmt"

func main() {
	age := 32 // Declaring and initializing 'age'

	// Working with pointers, a feature TypeScript doesn't have.
	var agePointer *int // Pointer to an int
	agePointer = &age   // Assigning the address of 'age' to agePointer

	fmt.Println("Age", *agePointer) // Dereferencing pointer to get the value of 'age'

	// Pass pointer to function to modify the value of 'age' indirectly
	editAgeAdultYears(agePointer)

	fmt.Println(age) // 'age' is modified by the function
}

// Function that takes a pointer to an int
func editAgeAdultYears(age *int) {
	*age = *age - 18 // Modifying the value at the pointer address
}
