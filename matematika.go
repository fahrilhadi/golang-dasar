package main

import "fmt"

func main()  {
	var a = 10
	var b = 20
	var c = 5
	var d = a + b - c

	fmt.Println(d)

	i := 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	j := 20
	j++
	fmt.Println(j)

}