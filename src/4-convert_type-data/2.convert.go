package main 

import (
	"fmt"
	"strconv"
)

func main() {
	gaji := "1000000"
	gajiConv, err := strconv.Atoi(gaji)
	if err != nil {
		fmt.Println("error convert a string...")
		return
	}

	// jika ingin mengabaikan err message, gunakan nama var _ untuk var err nya
	gajiConv2, _ := strconv.Atoi(gaji)
	bonus := gajiConv2 + 500000

	fmt.Println("gaji setelah terformat", gajiConv)
	fmt.Println("total bonus gaji : ", bonus)
}