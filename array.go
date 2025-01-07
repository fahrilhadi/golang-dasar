package main

import "fmt"

func main()  {
	// cara 1 : deklarasikan satu-satu
	var names [2]string

	names[0] = "Fahril"
	names[1] = "Hadi"

	fmt.Println(names[0])
	fmt.Println(names[1])

	// cara 2 : langsung dideklarasikan sekaligus array nya
	var values = [4]int{
		20,
		30,
		40,
		50,
	}

	fmt.Println(values)

	// cara 3 : gunakan ... bila tidak ingin menentukan data array awal
	nilai := [...]int{
		100,
		200,
		300,
		400,
	}

	fmt.Println(len(nilai))
	nilai[0] = 500
	fmt.Println(nilai)
}