package main

import "fmt"

func main()  {
	const firstName string = "Fahril"
	const lastName = "Hadi"

	fmt.Println(firstName)
	fmt.Println(lastName)

	// error, karena nilai constant tidak bisa diubah
	// firstName = "Hadi"
	// lastName = "Fahril"

	// cara 2 : deklarasi const dengan banyak sekaligus
	const (
		namaAwal string = "Fahril"
		namaAkhir string = "Hadi"
	)

	fmt.Println(namaAwal)
	fmt.Println(namaAkhir)
}