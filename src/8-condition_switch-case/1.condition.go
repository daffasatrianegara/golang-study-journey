package main

import (
	"fmt"
)

func checkHari(hari int) (hasil string) {
	switch {
	case hari == 1 :
		hasil = "senin"

	case hari == 2 :
		hasil = "selasa"

	case hari == 3 :
		hasil = "rabu"

	case hari == 4 :
		hasil = "kamis"

	case hari == 5 :
		hasil = "jumat"

	case hari == 6 :
		hasil = "sabtu"

	case hari == 7 :
		hasil = "minggu"

	// same as else in if else condition
	default :
		hasil = "tidak terdefinisi"
	}

	return
} 

func main() {
	hari := 10
	cek := checkHari(hari)
	fmt.Printf("hari ini adalah hari %v \n", cek)
}