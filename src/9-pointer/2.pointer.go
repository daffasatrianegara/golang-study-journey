package main 

import (
	"fmt"
)

func changePoint(point *int) {
	*point = 200
}

func main() {
	point := 100
	changePoint(&point)

	fmt.Println(point)
}