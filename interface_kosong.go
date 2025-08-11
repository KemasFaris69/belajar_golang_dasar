package belajar_golang_dasar

import "fmt"

func Ups() interface{} {
	return 1
	// return true
	// return "Ups"
}

func InterfaceKosong() {
	kosong := Ups()
	fmt.Println(kosong)
}
