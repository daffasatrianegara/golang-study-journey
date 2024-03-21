package main

import "fmt"

func main() {
	arrNum := [...]string {"muhammad", "daffa", "satria", "negara"}
	
	i := 0
	for i < len(arrNum) {
		fmt.Println(arrNum[i])
		i++
	}

	fmt.Println()
	
	for j := 0; j < len(arrNum); j++ {
		arrNum[j] = "kosong"
		fmt.Println(arrNum[j])
	}
}