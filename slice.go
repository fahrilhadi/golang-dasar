package main

import "fmt"

func main()  {
	names := [...]string{
		"Fahril",
		"Hadi",
		"Fadli",
		"Ramdani",
		"Abu",
		"Hanif",
	}

	slice1 := names[4:6]

	fmt.Println(slice1)

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	daySlice1 := days[5:]
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Fahril"
	newSlice[1] = "Hadi"
	// newSlice[2] = "Fadli" // error karena array ditetapkan menjadi 2 (tidak boleh lebih)

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "Sukri", "Joni")
	fmt.Println(newSlice2)

	newSlice2[0] = "Doni"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)
}