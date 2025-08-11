package belajar_golang_dasar

import "fmt"

func array() {
	// var names [3]string
	// names[0] = "eko"
	// names[1] = "kurniawan"
	// names[2] = "khannedy"

	// fmt.Println(names[0])
	// fmt.Println(names[1])
	// fmt.Println(names[2])

	//buat array secara langsung
	var values = [...]int{
		90,
		80,
		95,
		100,
		120,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	fmt.Println(values)

}
