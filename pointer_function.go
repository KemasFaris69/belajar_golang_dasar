package belajar_golang_dasar

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountyToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func pointerFunc() {
	var address *Address = &Address{}
	ChangeCountyToIndonesia(address)

	fmt.Println(address)
}
