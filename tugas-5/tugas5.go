package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	
	// soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"


	// soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"


	// soal 4
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(title string, jam string, genre string, tahun string) {
		var mapFilm = map[string]string{
			"genre": genre,
			"jam":   jam,
			"tahun": tahun,
			"title": title,
		}

		dataFilm = append(dataFilm, mapFilm)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, v := range dataFilm {
		fmt.Println(v)
	}

}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

func introduce(nama string, gender string, pekerjaan string, usia string) (rv string) {
	var pronouns string

	if gender == "laki-laki" {
		pronouns = "Pak"
	} else if gender == "perempuan" {
		pronouns = "Bu"
	}

	rv = pronouns + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun"
	return
}

func buahFavorit(nama string, buah ...string) (rv string) {
	var buahStr string

	for i := 0; i < len(buah); i++ {
		if i == len(buah)-1 {
			buahStr += `"` + buah[i] + `"`
		} else {
			buahStr += `"` + buah[i] + `", `
		}
	}

	rv = "halo nama saya " + strings.ToLower(nama) + " dan buah favorit saya adalah " + buahStr

	return
}
