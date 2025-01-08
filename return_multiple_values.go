package main

import "fmt"

func getFullName() (string, string)  {
	return "Fahril", "Hadi"
}

func main()  {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)
}