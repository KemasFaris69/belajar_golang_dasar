package belajar_golang_dasar

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application")
}

func run() {
	runApplication()
}
