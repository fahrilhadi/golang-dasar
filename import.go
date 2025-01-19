package main

import (
	"fmt"

	"github.com/fahrilhadi/golang-dasar/helper"
)

func main()  {
	result := helper.SayHello("Fahril")
	fmt.Println(result)

	fmt.Println(helper.version) // tidak bisa diakses dari luar package
	fmt.Println(helper.sayGoodBye("Fahril")) // tidak bisa diakses dari luar package
}