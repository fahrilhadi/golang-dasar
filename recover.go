package main

import (
	"fmt"
)

// kode program recover yang benar

func endApp()  {
	fmt.Println("Akhir Program")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

// kode program recover yang salah

/* func runApp(error bool)  {
	defer endApp()
	
	if error {
		panic("Error")
	}

	message := recover()
	fmt.Println("Terjadi Panic", message)
} */

func runApp(error bool)  {
	defer endApp()
	
	if error {
		panic("Error")
	}
}

func main()  {
	runApp(true)
	fmt.Println("Fahril Hadi")
}