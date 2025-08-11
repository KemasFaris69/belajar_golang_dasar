package belajar_golang_dasar

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}

func funcreturn() {
	result := getHello("eko")
	fmt.Println(result)

	fmt.Println(getHello("budi"))
	fmt.Println(getHello("cecep"))
}
