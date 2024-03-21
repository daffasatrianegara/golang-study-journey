package main

import (
	"fmt"
)


func checkAddress(num *int) {
	address := num
	fmt.Println(&address)
}

func main() {
	num := 100
	checkAddress(&num)
}