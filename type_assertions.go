package main

import (
	"fmt"
)

func random() any {
	return 100
}

func main()  {
	result := random()
	// resultString := result.(string)
	// fmt.Println(resultString)

	// jangan paksa menjadi result integer (panic)

	// resultInt := result.(int)
	// fmt.Println(resultInt)

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}