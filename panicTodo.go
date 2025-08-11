package belajar_golang_dasar

import "fmt"

func endApp() {
	fmt.Println("End app")
	message := recover()
	fmt.Println("terjadi panic", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Ups Error")
	}
}

func panicTodo() {
	runApp(true)
	fmt.Println("Eko Kurniawan Khannedy")
}
