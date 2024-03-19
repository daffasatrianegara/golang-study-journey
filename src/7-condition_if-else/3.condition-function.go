package main

import (
	"fmt"
)

func bayarUtang(uang float64, utang float64) {
	var total float64
	if uang > utang {
		total = uang - utang
		fmt.Printf("berhasil bayar hutang.\nsisa uang anda sebesar %v rupiah.", total)
	} else {
		total = utang - uang
		fmt.Printf("gagal bayar hutang.\nuang anda kurang %v rupiah.", total)
	}
}
func main() {
	bayarUtang(100000, 500000)
}