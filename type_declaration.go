package belajar_golang_dasar

import "fmt"

func typesdeclaration() {
	//cara 1
	// type noKTP string

	// var ktpEko noKTP = "1111111111"
	// fmt.Println(ktpEko)
	// fmt.Println(noKTP("2222222222"))

	//cara 2
	type noKTP string

	var ktpEko noKTP = "111111111"

	var contoh string = "21222222"
	var contohktp noKTP = noKTP(contoh)

	fmt.Println(ktpEko)
	fmt.Println(contohktp)
}
