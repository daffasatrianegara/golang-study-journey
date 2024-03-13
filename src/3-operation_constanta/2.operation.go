package main

import (
	"fmt"
)

func main() {
	a, b := 10, 5
	a++
	b--
	fmt.Printf("1 hasil a = %d dan hasil b = %d \n", a, b)
	a++
	b--
	fmt.Printf("2 hasil a = %d dan hasil b = %d \n", a, b)
	a++
	b--
	fmt.Printf("3 hasil a = %d dan hasil b = %d \n", a, b)
}