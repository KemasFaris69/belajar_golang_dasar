package belajar_golang_dasar

import "fmt"

func main() {
	name := "jocko"

	if name == "eko" {
		fmt.Println("Hello Eko")
	} else if name == "joko" {
		fmt.Println("Hello joko")
	} else if name == "budi" {
		fmt.Println("Hello budi")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}
	//short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
