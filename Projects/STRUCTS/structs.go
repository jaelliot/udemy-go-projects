package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createAt  time.Time
}

func (u *user) outputUserDetails() {
	//...

	fmt.Println(u.firstName, u.lastName, u.birthdate, u.createAt)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
	u.birthdate = ""
	u.createAt = time.Time{}
}

func newUser(firstName, lastName, birthdate string) user {
	return user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createAt:  time.Now(),
	}
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = newUser(userFirstName, userLastName, userBirthdate)

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
