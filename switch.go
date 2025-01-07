package main

import "fmt"

func main()  {
	name := "Abu Hanif"

	switch name {
	case "Fahril":
		fmt.Println("Halo, Fahril")
	case "Hadi":
		fmt.Println("Halo, Hadi")
	case "Fahril Hadi":
		fmt.Println("Halo, Fahril Hadi")
	default:
		fmt.Println("Halo, boleh kenalan?")
	}
}