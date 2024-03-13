package main

import (
	"fmt"
)

func main() {
	var pertambahan int
	a, b := 30, 20
	pertambahan = a + b 
	pengurangan := a - b
	fmt.Println("pertambahan 30 + 20 adalah :", pertambahan)
	fmt.Printf("pengurangan %d + %d adalah : %d \n", a, b, pengurangan)
}