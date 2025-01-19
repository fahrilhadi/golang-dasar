package accessmodifier

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "Good Bye " + name
}

func main()  {
	fmt.Println(version) // bisa diakses karena masih dalam package yang sama
}