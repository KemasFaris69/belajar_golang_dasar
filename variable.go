package belajar_golang_dasar

import "fmt"

func variables() {

	//cara pertama
	var nama string

	nama = "Faris"
	fmt.Println(nama)

	//cara kedua
	name := "Eko Kurniawan"
	name = "Eko Kurniawan"
	fmt.Println(name)

	name = "Eko Khannedy"
	fmt.Println(name)

	//cara ketiga
	var (
		firstname  = "Eko"
		middlename = "kurniawan"
		lastname   = "khannedy"
	)
	fmt.Println(firstname)
	fmt.Println(middlename)
	fmt.Println(lastname)
}
