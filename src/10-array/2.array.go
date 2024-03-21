package main

import (
	"fmt"
)

func main() {
	// change value in array
	// jika ingin menggunakan ... harus ada value terlebih dahulu
	var arrNum = [...]int {10, 20, 30, 40, 50}
	arrNum[4] = 100
	fmt.Println(arrNum)
}