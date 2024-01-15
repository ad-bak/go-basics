package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":   "https://google.com",
		"facebook": "https://facebook.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["google"])
	websites["Linkedin"] = "https://linkedin.com"

	fmt.Println(websites)

	delete(websites, "google")

	fmt.Println(websites)

}
