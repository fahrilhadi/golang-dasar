package main

import "fmt"

func main()  {
	type NoKTP string

	var ktpFahril NoKTP = "1111111111111"

	var contoh string = "2222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpFahril)
	fmt.Println(contohKtp)
}