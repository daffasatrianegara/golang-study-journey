package main

import (
	"fmt"
)

func isEvenOrOdd(nums ...int) (even []int, odd []int) {
	for _, num := range nums {
		if num%2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return
}

func main() {
	even, odd := isEvenOrOdd(10, 20, 100, 5, 7, 5, 3, 2, 100)
	fmt.Printf("even value\t: %v \nodd value\t: %v \n", even, odd)
}