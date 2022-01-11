package main

import (
	"fmt"
	"math"
)

func main() {
	// soal 1
	var luasLigkaran float64
	var kelilingLingkaran float64

	hitungLuasLingkaran(&luasLigkaran, 10)
	hitungKelilingLingkaran(&kelilingLingkaran, 10)

	fmt.Printf("%.2f\n", luasLigkaran)
	fmt.Printf("%.2f\n", kelilingLingkaran)

	// soal 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// soal 3
	var buah = []string{}

	tambahBuah(&buah)

	for i, item := range buah {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm(&dataFilm, "LOTR", "2 jam", "action", "1999")
	tambahDataFilm(&dataFilm, "avenger", "2 jam", "action", "2019")
	tambahDataFilm(&dataFilm, "spiderman", "2 jam", "action", "2004")
	tambahDataFilm(&dataFilm, "juon", "2 jam", "horror", "2004")

	for index, value := range dataFilm {
		var i int = 0
		for k, v := range value {
			if i == 0 {
				fmt.Printf("%d. %s : %s\n", index+1, k, v)
			} else {
				fmt.Printf("   %s : %s\n", k, v)
			}

			i++
		}
		fmt.Printf("\n")
	}
}

func hitungLuasLingkaran(luas *float64, r float64) {
	*luas = math.Pi * math.Pow(r, 2)
}

func hitungKelilingLingkaran(keliling *float64, r float64) {
	*keliling = 2 * math.Pi * r
}

func introduce(sentence *string, nama string, gender string, pekerjaan string, usia string) {
	var pronouns string

	if gender == "laki-laki" {
		pronouns = "Pak"
	} else if gender == "perempuan" {
		pronouns = "Bu"
	}

	*sentence = pronouns + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
}

func tambahBuah(buah *[]string) {
	*buah = append(*buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

func tambahDataFilm(dataFilm *[]map[string]string, title string, duration string, genre string, year string) {
	*dataFilm = append(*dataFilm, map[string]string{
		"genre":    genre,
		"duration": duration,
		"year":     year,
		"title":    title,
	})
}
