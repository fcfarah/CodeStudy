package main

import (
	"fmt"
	"time"
)

type user struct {
	FirstName string
	LastName  string
	Birthdate string // YYYY-MM-DD
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthdate,
		createdAt: time.Now(),
	}
	// appUser.FirstName = firstName
	// appUser.LastName = lastName
	// appUser.Birthdate = birthdate

	// ... do something awesome with that gathered data!

	fmt.Println(appUser)
	// fmt.Println(fappUser

	//

	outputUserDetails(appUser)
	// outputUserDetails(appUser.FirstName, appUser.LastName, appUser.Birthdate)
	// outputUserDetails(firstName, lastName, birthdate)
}

func outputUserDetails(u user) {
	fmt.Println("First Name: ", u.FirstName)
	fmt.Println("Last Name: ", u.LastName)
	fmt.Println("Birthdate: ", u.Birthdate)
	fmt.Println("Created At: ", u.createdAt)

	// fmt.Println("First Name: ", firstName)
	// fmt.Println("Last Name: ", lastName)
	// fmt.Println("Birthdate: ", birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
