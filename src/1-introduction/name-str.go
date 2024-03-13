package main

import (
	"fmt"
)

func main() {
	fmt.Print("halooo, nama saya daffa, ")
	fmt.Println("umur saya 21.")
	
	// wajib pake printf
	fmt.Printf("hobi saya adalah %s, %s, dan %s \n", "bermain sepak bola", "membaca buku", "merokok")
	// %s means to String, if your number are decimal use %d
	fmt.Printf("saya berumur %d", 21)
}