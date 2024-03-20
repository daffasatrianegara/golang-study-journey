package main

import (
	"fmt"
)

func main() {
	// & for check point address in memory
	// * for check value in point address
	var point = 100
	var pointAdress *int = &point

	// change value 
	*pointAdress = 200

	// value
	fmt.Println(*pointAdress)
	fmt.Println(point)

	// address
	fmt.Println()
	fmt.Println(pointAdress)
	fmt.Println(&point)
}