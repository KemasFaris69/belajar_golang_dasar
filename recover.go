package belajar_golang_dasar

import "fmt"

// contoh recover yang salah
func RunApp(error bool) {
	defer endApp()
	if error {
		panic("Ups ERROR")
	}
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func Recovers() {
	runApp(true)
	fmt.Println("Eko kurniawan Khannedy")
}
