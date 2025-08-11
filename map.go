package belajar_golang_dasar

import "fmt"

func mapFunc() {

	book := make(map[string]string)
	book["Tittle"] = "buku go-lang"
	book["Author"] = "Eko kurniawan"
	book["wrong"] = "Ups"

	delete(book, "wrong")
	fmt.Println(book)
	// person := map[string]string{
	// 	"name":    "acen",
	// 	"address": "cibuntu",
	// }

	// fmt.Println(person)
	// fmt.Println(person["name"])
	// fmt.Println(person["address"])

	//versi cepatnya
	// person := map[string]string{
	// 	"name":    "nugraha",
	// 	"address": "TG",
	// }
	// fmt.Println(person)
}
