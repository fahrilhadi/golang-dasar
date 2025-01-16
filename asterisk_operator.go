package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main()  {
	address1 := Address{"Pekanbaru", "Riau", "Indonesia"}
	address2 := &address1
	address2.City = "Dumai"
	fmt.Println(address1) // ikut berubah
	fmt.Println(address2) // berubah menjadi Dumai

	*address2 = Address{"Jakarta", "DKI Jakarta", "Indonesia"} // operator asterisk
	fmt.Println(address1)
	fmt.Println(address2)

}