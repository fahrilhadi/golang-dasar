package main

import "fmt"

type Costumer struct {
	Name 		string
	Address 	string
	Age			int
}

func main()  {
	// membuat data struct
	var fahril Costumer
	fahril.Name = "Fahril Hadi"
	fahril.Address = "Pekanbaru"
	fahril.Age = 24

	fmt.Println(fahril)
	fmt.Println(fahril.Name)
	fmt.Println(fahril.Address)
	fmt.Println(fahril.Age)

	// struct literals
	fadli := Costumer {
		Name: "Fadli",
		Address: "Siak",
		Age: 20,
	}

	fmt.Println(fadli)

	// atau
	hanif := Costumer {"Hanif", "Pelalawan", 29}
	fmt.Println(hanif)
}