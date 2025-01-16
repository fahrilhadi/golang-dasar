package main

import "fmt"

type Address struct{
	City, Province, Country string
}

func main()  {
	// address1 := Address{"Pekanbaru", "Riau", "Indonesia"}
	// address2 := address1 // di copy value nya (pass by value)

	// address2.City = "Dumai"
	// fmt.Println(address1) // tidak berubah city nya
	// fmt.Println(address2) // berubah city nya

	address1 := Address{"Pekanbaru", "Riau", "Indonesia"}
	address2 := &address1 // pass by reference (pointer)

	address2.City = "Dumai"
	fmt.Println(address1) // berubah city nya
	fmt.Println(address2) // berubah city nya
}