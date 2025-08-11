package belajar_golang_dasar

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter Filter) {
	filteredName := filter(name)
	fmt.Println("Hello", filteredName)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func funcAsParam() {
	sayHelloWithFilter("Anjing", spamFilter)

	// filter := spamFilter
	// sayHelloWithFilter("Anjing", filter)
}
