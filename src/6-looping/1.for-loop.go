package main

import (
	"fmt"
)

func main() {
	arrStr := []string {"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan", "sembilan", "sepuluh"}
	for i := 1; i <= 10; i++ {
		fmt.Printf("angka %v bentuk desimalnya adalah\t: %d \n", arrStr[i-1], i)
	}
}