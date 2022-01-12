package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var bangunDatar hitungBangunDatar = persegiPanjang{panjang: 10, lebar: 7}
	fmt.Printf("Luas Persegi Panjang : %d\n", bangunDatar.luas())
	fmt.Printf("Keliling Persegi Panjang : %d\n", bangunDatar.keliling())

	bangunDatar = segitigaSamaSisi{alas: 6, tinggi: 5}
	fmt.Printf("Luas Segitiga : %d\n", bangunDatar.luas())
	fmt.Printf("Keliling Segitiga : %d\n", bangunDatar.keliling())

	var bangunRuang hitungBangunRuang = tabung{jariJari: 7, tinggi: 18}
	fmt.Printf("Luas Permukaan Tabung : %.2f\n", bangunRuang.luasPermukaan())
	fmt.Printf("Volume Tabung : %.2f\n", bangunRuang.volume())

	bangunRuang = balok{panjang: 28, lebar: 24, tinggi: 10}
	fmt.Printf("Luas Permukaan Balok : %.2f\n", bangunRuang.luasPermukaan())
	fmt.Printf("Volume Balok : %.2f\n", bangunRuang.volume())

	fmt.Println("==================")

	// soal 2
	var ph phoneInterface = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Printf("%s", ph.display())

	fmt.Println("==================")

	// soal 3
	fmt.Println(luasPersegi(4, true))

	fmt.Println(luasPersegi(8, false))

	fmt.Println(luasPersegi(0, true))

	fmt.Println(luasPersegi(0, false))

	fmt.Println("==================")

	// soal 4
	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var sum int
	var numbers []string

	for _, v := range kumpulanAngkaPertama.([]int) {
		sum += v
		numbers = append(numbers, strconv.Itoa(v))
	}

	for _, v := range kumpulanAngkaKedua.([]int) {
		sum += v
		numbers = append(numbers, strconv.Itoa(v))
	}

	fmt.Println(prefix.(string) + strings.Join(numbers, " + ") + " = " + strconv.Itoa(sum))

}

// Segitiga Sama Sisi
type segitigaSamaSisi struct {
	alas, tinggi int
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

// Persegi Panjang
type persegiPanjang struct {
	panjang, lebar int
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

// Tabung
type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * math.Pow(t.jariJari, 2) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	var luasTutup float64 = 2 * math.Pi * math.Pow(t.jariJari, 2)
	var panjangSelimut float64 = 2 * math.Pi * t.jariJari
	var luasSelimut float64 = panjangSelimut * t.tinggi

	return luasTutup + luasSelimut
}

// Balok
type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) luasPermukaan() float64 {
	return float64(2 * ((b.panjang * b.lebar) + (b.panjang * b.tinggi) + (b.lebar * b.tinggi)))
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

// Interfaces
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

type phoneInterface interface {
	display() string
}

func (ph phone) display() string {
	return "name : " + ph.name + "\n" + "brand : " + ph.brand + "\n" + "year : " + strconv.Itoa(ph.year) + "\n" + "colors : " + strings.Join(ph.colors, ", ") + "\n"
}

func luasPersegi(sisi int, x bool) interface{} {
	var luas int = sisi * sisi

	if sisi != 0 {
		if x {
			return "luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(luas) + " cm"
		}

		return sisi
	} else {
		if x {
			return "Maaf anda belum menginput sisi persegi"
		}
	}

	return nil
}
