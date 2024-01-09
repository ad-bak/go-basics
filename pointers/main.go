package main

import "fmt"

func main() {
	x := 20
	y := 2

	add(&x, &y) // Passing the address of x

	fmt.Println(x) // x will be modified if add changes the value pointed to by the pointer
	fmt.Println(y) // y remains unchanged
}

func add(a *int, b *int) {
	*a = *a + *b // Dereferencing a to get/set its value
}
