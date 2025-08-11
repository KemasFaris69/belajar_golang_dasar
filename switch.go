package belajar_golang_dasar

import "fmt"

func SwitchFunc() {
	name := "joko"

	switch name {
	case "eko":
		fmt.Println("Hello eko")
	case "budi":
		fmt.Println("Hello budi")
	default:
		fmt.Println("hi, boleh kenalan?")
	}

	//short statement

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	name = "Khannedy"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}
}
