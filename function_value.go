package belajar_golang_dasar

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func funcValues() {
	goodbye := getGoodBye
	fmt.Println(goodbye("eko"))
}
