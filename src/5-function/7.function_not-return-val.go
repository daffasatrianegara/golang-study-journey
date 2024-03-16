package main

import (
	"fmt"
)

var text = "variabel global"

func printText() {
	fmt.Println("ini " + text)
}

func main() {
	printText()
}