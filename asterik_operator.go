package belajar_golang_dasar


import "fmt"

type address struct {
	City, Province, Country string
}

func asterik() {
	var Address1 Address = Address{"subang", "jawa barat", "Indonesia"}
	var Address2 *Address = &Address1 // pointer

	Address2.City = "Bandung"
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2) // berubah menjadi bandung

	// Address2 = &Address{"jakarta", "dki jakarta", "indonesia"}
	*Address2 = Address{"jakarta", "dki jakarta", "indonesia"}
	fmt.Println(Address1)
	fmt.Println(Address2)
}
