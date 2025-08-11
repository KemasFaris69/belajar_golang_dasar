package belajar_golang_dasar

import "fmt"

func Fortodo() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("perulangan ke :", counter)
	// 	counter++

	for counter := 1; counter <= 10; counter++ {
		fmt.Println("perulangan ke :", counter)

	}
	fmt.Println("Selesai")

	names := []string{"eko", "kurniawan", "Khannedy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, names := range names {
		fmt.Println("Index", index, "=", names)
	}

	//jika tidak butuh index
	for _, names := range names {
		fmt.Println(names)
	}
}
