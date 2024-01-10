package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func newUser(firstName, lastName, birthDate string) (*user, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate values are required!")
	}

	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func (u *user) outputUserDetails() {

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.outputUserDetails()

	appUser.clearUserName()

	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
