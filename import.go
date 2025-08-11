package belajar_golang_dasar

import (
	"belajar_golang_dasar/helper"
	"fmt"
)

func ImportFunc() {
	result := helper.SayHello("Eko")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // tidak bisa diakses
	// fmt.Println(helper.sayGoodBye("Eko")) // tidak bisa diakses
}
