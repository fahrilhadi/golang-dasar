package main

import "fmt"

type Customer struct {
	Name		string
	Address		string
	Age			int
}

func (customer Customer) sayHello(name string)  {
	fmt.Println("Halo", name, "nama saya", customer.Name)
}

func main()  {
	fahril := Customer{Name: "Fahril"}
	fahril.sayHello("Fadli")
}