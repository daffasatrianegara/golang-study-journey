package main

import (
	"fmt"
)

var uang, utang int
var bayar float64

func main() {
	uang = 500000
	utang = 500000
	if uang > utang {
		bayar = float64(uang) - float64(utang)
		fmt.Printf("bisa bayar hutang \nsekarang uang anda sebesar %v rupiah", bayar)
	} else if uang == utang {
		bayar = float64(uang) / 2
		fmt.Printf("bayar hutang setengah dulu \nsebesar %v rupiah", bayar)
	} else {
		fmt.Println("belum bisa bayar hutang")
	}
}