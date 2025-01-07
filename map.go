package main

import "fmt"

func main()  {
	person := map[string]string {
		"name" : "Fahril Hadi",
		"address" : "Pekanbaru",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Golang"
	book["author"] = "Fahril Hadi"
	book["wrong"] = "Ups"

	fmt.Println(book)

	delete(book, "wrong")
	fmt.Println(book)
}