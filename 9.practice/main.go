package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	//1
	hobbies := [3]string{"Programming", "Football", "Reading"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	//3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	//4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	//5
	goals := []string{"Learn to code", "Be good at it"}
	fmt.Println(goals)

	//6
	goals[1] = "Be great at it"
	goals = append(goals, "Keep coding")
	fmt.Println(goals)

	products := []Product{
		{
			title: "VR Headset",
			id:    "1",
			price: 3500.0,
		},
		{
			title: "Laptop",
			id:    "2",
			price: 1200.0,
		},
	}

	fmt.Println(products)

}
