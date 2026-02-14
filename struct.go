package main

import "fmt"

func structFunc() {
	// fmt.Println("Hello From Struct File")

	firstName := getUserData("Enter your First name: ")
	lastName := getUserData("Enter your Last name: ")

	fmt.Println("Your full name is: ", firstName+" "+lastName)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string
	fmt.Scan(&value)

	return value

}
