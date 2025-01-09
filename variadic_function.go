package main

import "fmt"

func sumAll(numbers ...int) int  {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func main()  {
	fmt.Println(sumAll(10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10))
	fmt.Println(sumAll(10, 10, 10, 10, 10, 10, 10, 10, 10))

	// kasus dimana kita menggunakan variadic function, namun memiliki variable berupa slice
	// kita bisa menjadikan slice sebagai vararg parameter

	numbers := []int{10, 10, 10, 10}
	fmt.Println(sumAll(numbers...))
}