package main

import (
	"fmt"
)

func main() {
	// array dapat dideklarasikan melalui 2 cara :
	arrMobil := [3]string {}
	var arrMobil2 [3]string
	arrMobil[0] = "Xenia"
	arrMobil[1] = "Avanza"
	arrMobil[2] = "BMW"

	for i := 0; i < len(arrMobil); i++ {
		arrMobil2[i] = arrMobil[i]
	}

	fmt.Printf("metode 1\t: %v\nmetode 2\t: %v", arrMobil, arrMobil2)
}