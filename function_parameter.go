package belajar_golang_dasar

import "fmt"

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func FuncParameter() {
	sayHelloTo("Eko", "Khannedy")
	sayHelloTo("Acen", "Nugraha")
}
