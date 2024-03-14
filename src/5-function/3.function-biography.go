package main

import (
	"fmt"
	"strconv"
)

func getBiography(nama string, pekerjaan string, umur int) string {
	umurConv := strconv.Itoa(umur)
	return nama + " memiliki pekerjaan " + pekerjaan + ", dan berumur " + umurConv + " tahun"
}

func main() {
	fmt.Println(getBiography("Daffa", "Fullstack web developer", 21))
}