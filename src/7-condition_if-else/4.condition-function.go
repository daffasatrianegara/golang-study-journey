package main

import (
	"fmt"
	"math/rand"
)

func GKB(play int) (hasil string, playerAnswer string, computerAnswer string) {
	GKB := [3]string {"Gunting", "Kertas", "Batu"}
	randomAnsw := rand.Intn(len(GKB))
	playerAnswer = GKB[play-1]
	computerAnswer = GKB[randomAnsw]
	if (playerAnswer == "Gunting" && computerAnswer == "Kertas") || (playerAnswer == "Kertas" && computerAnswer == "Batu") || (playerAnswer == "Batu" && computerAnswer == "Gunting") {
		hasil = "Selamat!, anda menang!!!"
	} else if (playerAnswer == "Gunting" && computerAnswer == "Batu") || (playerAnswer == "Kertas" && computerAnswer == "Gunting") || (playerAnswer == "Batu" && computerAnswer == "Kertas") {
		hasil = "Anda kalah, coba lagi..."
	} else if (playerAnswer == computerAnswer) {
		hasil = "Pertandingan seri"
	} else {
		hasil = "ERROR, nilai tidak ada"
	}
	return
}

func main() {
	jawaban := 4
	hasil, playerAnswer, computerAnswer := GKB(jawaban)
	fmt.Printf("\nPERMAINAN GUNTING, KERTAS, BATU\nAnda memilih\t\t: %v\nKomputer memilih\t: %v\nHASIL\t\t\t: %v\n\n", playerAnswer, computerAnswer, hasil)
}