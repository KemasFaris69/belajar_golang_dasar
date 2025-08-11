package belajar_golang_dasar

import "fmt"

func random() any {
	return "OK"

}

func typeAssertionFunc() {
	var result any = random()
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
