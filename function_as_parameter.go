package main

import "fmt"

// jika function terlalu panjang, maka type declaration bisa digunakan untuk membuat alias function

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter)  {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string  {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main()  {
	sayHelloWithFilter("Fahril Hadi", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("Anjing", filter)
}