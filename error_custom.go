package belajar_golang_dasar

import (
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{Message: "Validation Error"}
	}

	if id != "eko" {
		return &validationError{Message: "data not found"}
	}

	return nil
}

func erorcustom() {
	err := SaveData("eko", nil)
	if err != nil {
		// if validationErr, ok := err.(*validationError); ok {
		// 	fmt.Println("validation Error :", validationErr.Message)
		// } else if notFoundErr, ok := err.(*notFoundError); ok {
		// 	fmt.Println("not found error:", notFoundErr.Message)
		// } else {
		// 	fmt.Println("Unknown error:", err.Error())
		// }
		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error:", finalError.Error())
		case *notFoundError:
			fmt.Println("not found error:", finalError.Error())
		default:
			fmt.Println("Unknown error:", err.Error())
		}

	} else {
		fmt.Println("Sukses")
	}
}
