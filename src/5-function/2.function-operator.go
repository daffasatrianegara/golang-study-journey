package main 

import (
	"fmt"
	"strconv"
)

func convertInt(num string) int {
	convNum, _ := strconv.Atoi(num)
	return convNum
}

func pertambahan(num1 int, num2 int) int {
	return num1 + num2
}

func pengurangan(num1 int, num2 int) int {
	return num1 - num2
}

func main() {
	gaji := "100000"
	convGaji := convertInt(gaji)
	fmt.Println(pertambahan(10, 20))
	fmt.Println(pengurangan(convGaji, 50000)) // 50000
}