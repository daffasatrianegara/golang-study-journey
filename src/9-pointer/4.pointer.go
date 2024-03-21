package main

import (
	"fmt"
)

func main() {
	number := 100
	numberVal := &number

	fmt.Println(&number)
	fmt.Println(*numberVal)
}