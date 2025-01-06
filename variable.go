package main

import "fmt"

func main()  {
	// cara 1 : tipe data disertakan
	var name string

	name = "Fahril"
	fmt.Println(name)

	name = "Hadi"
	fmt.Println(name)

	// cara 2 : tipe data tidak disertakan atau langsung diinisialisasikan
	var nama = "Fahril"
	fmt.Println(nama)

	nama = "Hadi"
	fmt.Println(nama)

	// cara 3 : deklarasi variabel menggunakan :=
	namaa := "Fahril"
	fmt.Println(namaa)

	namaa = "Hadi"
	fmt.Println(namaa)

	// cara 4 : variabel lebih dari 1
	var (
		firstName = "Fahril"
		lastName = "Hadi"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}