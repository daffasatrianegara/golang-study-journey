package main

import "fmt"

func main() {
	i := 1
	end := 10
	var sum int
	for i <= end {
		sum += i
		fmt.Println(sum)
		i++
	}

	fmt.Println("jumlah angka sum setelah penjumlahan selesai adalah : ", sum)
}