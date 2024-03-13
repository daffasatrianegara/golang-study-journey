package main

import (
	"fmt"
)

func main() {
	// tidak bisa berubah
	const a = 100
	//a++ //ERROR
	fmt.Println(a)
	fmt.Println(a + a)
	fmt.Println(a - a)
}