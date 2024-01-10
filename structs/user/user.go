package user

import (
	"errors"
	"fmt"
	"time"
)

// User struct defines the properties of a user
type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Admin struct, embedding User, adds additional properties for an admin
type Admin struct {
	email    string
	password string
	User
}

// NewAdmin creates and returns a new Admin with specified email and password.
// It also initializes the embedded User with default values.
func NewAdmin(email, password string) Admin {
	return Admin{
		email:    email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthDate: "___",
			createdAt: time.Now(),
		},
	}
}

// New creates a new User instance. Returns an error if required fields are missing.
func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("First name, last name and birthdate values are required!")
	}

	// If validation passes, return a new User instance
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

// OutputUserDetails prints the user details to standard output.
func (u *User) OutputUserDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

// ClearUserName clears the user's first and last names.
func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}
