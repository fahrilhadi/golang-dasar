package main

import "fmt"

type Man struct {
	Name	string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main()  {
	fahril := Man{"Fahril"}
	fahril.Married()

	fmt.Println(fahril.Name)
}