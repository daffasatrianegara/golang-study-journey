package main  

import (
	"fmt"
	"strconv"
)

func ramalan(nama string, umur int) (string, int) {
	namaOrg, ramalan := "manusia bernama " + nama,  umur + 20
	return namaOrg, ramalan
}

func convString(num int) string {
	conv := strconv.Itoa(num)
	return conv
}

func main() {
	namaOrg, ramalan := ramalan("daffa", 21)
	hasilRamalan := "20 tahun lagi akan berumur " + convString(ramalan)
	fmt.Printf("%s, \n%s \n", namaOrg, hasilRamalan)
}