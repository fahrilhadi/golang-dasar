package main

import "fmt"

func getCompleteName() (firstName, lastName string)  {
	firstName = "Fahril"
	lastName = "Hadi"

	return firstName, lastName
}

func main()  {
	firstName, lastName := getCompleteName()
	fmt.Println(firstName, lastName)
}