package main

import (
	"fmt"
	"strconv"
)

func main() {
	gaji := 1000
	gaji2 := strconv.Itoa(gaji)
	
	// ERR : type data difference
	//fmt.Println("gaji saya adalah " + gaji)
	fmt.Println("gaji saya adalah ", gaji)

	// convert int to str
	fmt.Println("gaji saya adalah " + strconv.Itoa(gaji) + " || " + gaji2)
}