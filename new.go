package main

import "fmt"

type Alamat struct {
	Kota, Provinsi, Negara string
}

func main()  {
	alamat1 := new(Alamat) // untuk data nya kosong (gunakan new)
	alamat2 := alamat1

	alamat2.Negara = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}