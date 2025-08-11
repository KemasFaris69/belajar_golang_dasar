package belajar_golang_dasar

import "fmt"

func closures() {
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter++
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
