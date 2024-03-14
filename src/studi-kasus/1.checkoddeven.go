package main

import (
	"fmt"
)

func checkVal(num int) string {
	var result string 
	if num % 2 == 0 {
		result = "the value is even"
		} else {
		result = "the value is odd"
	}
	return result
}

func main() {
	checkType := checkVal(5)
	fmt.Println(checkType)
}