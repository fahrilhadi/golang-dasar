package main

import (
	"fmt"

	"github.com/fahrilhadi/golang-dasar/database"
	_ "github.com/fahrilhadi/golang-dasar/internal" // blank identifier
)

func main()  {
	fmt.Println(database.GetDatabase())
}