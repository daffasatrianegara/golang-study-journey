package main

import (
	"fmt"
)

func printNumber() int {
	return 10 * 10 // 100
}

func printStr() string {
	hello := "hello world"
	return hello
}

func Print() string {
	fmt.Println("saya ngeprint didalam function...")
	return "daffa"
}

func main() {
	fmt.Println(printNumber() / 10) // 10
	fmt.Println(printStr() + ", nama saya daffa")
	fmt.Println(Print())
}