package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	// Get user data by prompting the user
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User // Declare a pointer to a User struct
	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err) // Print the error if user creation fails
		return
	}

	// Creating an Admin user
	admin := user.NewAdmin("test@example.com", "test123")
	admin.OutputUserDetails() // Print details of the admin user
	admin.ClearUserName()     // Clear the name of the admin user
	admin.OutputUserDetails() // Print details of the admin user after clearing the name

	appUser.OutputUserDetails() // Print details of the app user
	appUser.ClearUserName()     // Clear the name of the app user
	appUser.OutputUserDetails() // Print details of the app user after clearing the name
}

// getUserData prompts the user with the given text and returns the input
func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
