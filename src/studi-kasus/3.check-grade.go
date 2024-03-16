package main

import (
	"fmt"
)

func checkGrade(value int) (grade string) {
	if value <= 100 && value >= 85 {
		grade = "A"
	} else if value < 85 && value >= 70 {
		grade = "B"
	} else if value < 70 && value >= 60 {
		grade = "C"
	} else if value < 60 && value >= 40 {
		grade = "D"
	} else if value < 40 && value >= 1 {
		grade = "E"
	} else if value == 0 {
		grade = "F"
	} else {
		fmt.Println("nilai tidak terdefinisi")
		grade = "null"
	}
	return grade
}

func main() {
	value := 0
	grade := checkGrade(value)
	fmt.Printf("grade dari nilai yang anda input adalah %v \n", grade)
	if grade == "D" || grade == "E" || grade == "F" {
		fmt.Println("tolong belajar lebih giat lagi. semangat!!!")
	}
}