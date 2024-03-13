package main

import (
	"fmt"
)

// variabel
var namaDepan, namaBelakang = "Muhammad Daffa", "Satria Negara"
var namaLengkap = namaDepan + " " + namaBelakang
var a, b, c int = 1, 2, 3

func main() {
	fmt.Printf("nama depan saya adalah %s \n", namaDepan)
	fmt.Printf("dan nama belakang saya adalah %s \n", namaBelakang)
	fmt.Println("saya memiliki nama lengkap", namaLengkap)
	
	pertambahan := a + b + c 
	fmt.Printf("%d + %d + %d = %d", a, b, c, pertambahan)

}