package main

import (
	"fmt"
)

func main() {
	// metode variabel ini hanya bisa digunakan dalam fungsi
	sekolah := "Peradaban"
	namaDepan, namaBelakang := "Daffa", "Satria"

	fmt.Printf("nama saya adalah %s %s, bisa dipanggil %s \n", namaDepan, namaBelakang, namaDepan)
	fmt.Println("saya bersekolah di", sekolah + ".")
}