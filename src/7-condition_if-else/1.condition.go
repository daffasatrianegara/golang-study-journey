package main 

import (
	"fmt"
)

func main() {
	i := 75

	if i <= 100 && i >= 80 {
		fmt.Println("Nilai anda besar, yaitu :", i)
	} else if i < 80 && i >= 60 {
		fmt.Println("nilai anda lumayan, yaitu :", i)
	} else if i < 60 && i >= 40 {
		fmt.Println("nilai anda cukup, yaitu :", i)
	} else if i < 40 && i >= 0 {
		fmt.Println("nilai anda buruk, yaitu :", i)
	} else {
		fmt.Println("nilai tidak terdefinisi...")
	}
}