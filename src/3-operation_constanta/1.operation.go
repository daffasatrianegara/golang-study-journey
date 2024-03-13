package main

import (
	"fmt"
)

var a, b,c, d, e int = 10, 20, 30, 40, 50

func main() {
	pertambahan := a + b
	pengurangan := c - e
	pembagian := e / a
	perkalian := a * a
	campuran := ((a+b)*(a-e))/a
	fmt.Printf("pertambahan = %d \npengurangan = %d \npembagian = %d \nperkalian = %d \ncampuran = %d \n", pertambahan, pengurangan, pembagian, perkalian, campuran)
}