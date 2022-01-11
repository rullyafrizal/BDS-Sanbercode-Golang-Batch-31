package main

import (
	"fmt"
	"math"
)

func main() {
	// soal 1
	var b1 buah = buah{"Nanas", "Kuning", false, 9000}
	var b2 buah = buah{"Jeruk", "Orange", true, 8000}

	var b3 buah
	b3.nama = "Semangka"
	b3.warna = "Hijau & Merah"
	b3.adaBijinya = true
	b3.harga = 10000

	var b4 buah = buah{nama: "Pisan", warna: "Kuning", adaBijinya: false, harga: 5000}

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
	fmt.Println(b4)

	// soal 2
	var s segitiga = segitiga{7, 10}
	fmt.Printf("Luas Segitiga : %d\n", s.luas())

	var p persegi = persegi{5}
	fmt.Printf("Luas Persegi : %d\n", p.luas())

	var pp persegiPanjang = persegiPanjang{15, 12}
	fmt.Printf("Luas Persegi Panjang : %d\n", pp.luas())

	// soal 3
	var ph = &phone{}
	ph.addColor("Kuning", "Hijau", "Merah")
	fmt.Println(ph.colors)

	// soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for idx, val := range dataFilm {
		fmt.Printf("%d. ", idx+1)

		if idx+1 == 3 {
			fmt.Printf("genre : %s\n", val.genre)
			fmt.Printf("   year : %d\n", val.year)
			fmt.Printf("   title : %s\n", val.title)
			fmt.Printf("   duration : %d\n", val.duration)
		} else {
			fmt.Printf("title : %s\n", val.title)
			fmt.Printf("   duration : %d\n", val.duration)
			fmt.Printf("   genre : %s\n", val.genre)
			fmt.Printf("   year : %d\n", val.year)
		}

		fmt.Printf("\n")
	}
}

type buah struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      uint32
}

type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luas() int {
	return s.alas * s.tinggi / 2
}

func (p persegi) luas() int {
	return int(math.Pow(float64(p.sisi), 2))
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (ph *phone) addColor(colors ...string) {
	for _, v := range colors {
		ph.colors = append(ph.colors, v)
	}
}

type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	var mov movie = movie{title, genre, duration, year}

	*dataFilm = append(*dataFilm, mov)
}
