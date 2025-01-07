package main

import "fmt"

func main()  {
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("Perulangan ke ", counter)
	}

	fmt.Println("Selesai")

	names := []string {"Fahril", "Hadi", "Fadli", "Ramdani"}
	for index, name := range names {
		fmt.Println("Index", index, "=", name)
	}
}