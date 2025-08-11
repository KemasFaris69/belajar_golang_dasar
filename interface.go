package belajar_golang_dasar

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func InterfaceFunc() {
	person := Person{Name: "Eko"}
	SayHello(person)

	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
