package belajar_golang_dasar

import (
	"belajar_golang_dasar/database"
	_ "belajar_golang_dasar/database"
	_ "belajar_golang_dasar/internal"
	"fmt"
)

func initFunc() {
	fmt.Println(database.GetDatabase())
}
