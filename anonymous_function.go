package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist)  {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main()  {
	// cara pertama anonymous function
	blacklist := func (name string) bool  {
		return name == "Anjing"
	}
	registerUser("Fahril Hadi", blacklist)

	// cara kedua anonymous function
	registerUser("Anjing", func(name string) bool {
		return name == "Anjing"
	})
}