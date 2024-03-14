package main

import (
	"fmt"
	"strconv"
)

func greeting(nama string, prodi string, hobi string, umur int, universitas string, alamat string) (string, string) {
	umurConv := strconv.Itoa(umur)
	return "halo, perkenalkan nama saya " + nama + ", berumur " + umurConv + ", dari prodi " + prodi,
	"saya memiliki hobi " + hobi + ", sekarang kuliah di " + universitas + ", buat temen-temen yang mau main ke tempat saya bisa ke alamat " + alamat
}

func main() {
	greeting1, greeting2 := greeting("daffa", "teknologi informasi", "bermain game", 21, "UNY", "kost al-faruq, condong catur")
	fmt.Printf("perkenalan pertama\t: %s \nperkenalan kedua\t: %s \n", greeting1, greeting2)
}