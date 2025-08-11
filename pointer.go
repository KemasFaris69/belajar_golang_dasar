package belajar_golang_dasar

import "fmt"

type Address1 struct {
	City, Province, Country string
}

func pointer() {
	// Address1 := Address{"subang", "jawa barat", "Indonesia"}
	// Address2 := Address1 // copy value

	// Address2.City = "Bandung"
	// fmt.Println(Address1) // tidak berubah
	// fmt.Println(Address2) // city berubah menjadi bandung

	var Address1 Address = Address{"subang", "jawa barat", "Indonesia"}
	var Address2 *Address = &Address1 // pointer

	Address2.City = "Bandung"
	fmt.Println(Address1) // ikut berubah
	fmt.Println(Address2) // berubah menjadi bandung
}
