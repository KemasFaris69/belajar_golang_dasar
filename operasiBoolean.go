package belajar_golang_dasar

import "fmt"

func boolFunc() {
	var nilaiAkhir = 90
	var absensi = 81

	//requirement
	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

}
