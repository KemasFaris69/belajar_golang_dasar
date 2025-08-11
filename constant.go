package belajar_golang_dasar


import "fmt"

func constant() {
	// cara pertama
	// const firstname = "eko"
	// const middlename = "kurniawan"
	// const lastname = "khannedy"

	// fmt.Println(firstname)
	// fmt.Println(middlename)
	// fmt.Println(lastname)

	//cara kedua multiple constant
	const (
		firstname = "Edi"
		lastname  = "nugraha"
	)
	fmt.Println(firstname)
	fmt.Println(lastname)
}
