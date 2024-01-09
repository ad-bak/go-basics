package main

import "fmt"

func main() {
	x := 20
	y := 2

	add(&x, y)

	fmt.Println(x)
	fmt.Println(y)

}

func add(a, b *int) int {
	return *a + *b
}
