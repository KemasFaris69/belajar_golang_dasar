package belajar_golang_dasar

import "fmt"

func getFullname() (string, string) {
	return "eko", "khannedy"
}

func multipleValue() {
	// firstName, lastName := getFullname()
	// fmt.Println(firstName, lastName)

	firstName, _ := getFullname()
	fmt.Println(firstName)
}
