package belajar_golang_dasar

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func pointerMethod() {
	eko := Man{"Eko"}
	eko.Married()
	fmt.Println(eko.Name)

}
