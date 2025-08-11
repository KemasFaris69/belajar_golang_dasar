package belajar_golang_dasar

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "eko"
	middleName = "kurniawan"
	lastName = "khannedy"

	return firstName, middleName, lastName
}

func namedReturnValue() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
