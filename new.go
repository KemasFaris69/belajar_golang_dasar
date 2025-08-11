package belajar_golang_dasar

import "fmt"

type addresses struct {
	City, Province, Country string
}

func newFunc() {
	var alamat1 *Address = &Address{}
	var alamat2 *Address = alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
