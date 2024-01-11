package main

import "fmt"

// we extend the string type to be able to create methods for it
type str string

// here we create a method for the str type
// just an example of how to create methods for types
// not really useful in this case
func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "John"

	name.log()
}
