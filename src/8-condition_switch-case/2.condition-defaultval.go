package main 

import (
	"fmt"
)

func lampuLalin(warna string) {
	switch warna {
	case "merah" :
		fmt.Println("Berhenti")

	case "kuning" :
		fmt.Println("Hati-hati")

	case "hijau" :
		fmt.Println("Jalan")

	default :
		fmt.Println("lampu lalu lintas error...")
	}
} 

func main() {
	warna := "hijau"
	lampuLalin(warna)
}