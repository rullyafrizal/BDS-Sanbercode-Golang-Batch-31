package main

import (
	"fmt"
	"strconv"
	"strings"
	. "tugas-9/lib"
)

func main() {
	// soal 1
	var bangunDatar HitungBangunDatar = PersegiPanjang{Panjang: 10, Lebar: 7}
	fmt.Printf("Luas Persegi Panjang : %d\n", bangunDatar.Luas())
	fmt.Printf("Keliling Persegi Panjang : %d\n", bangunDatar.Keliling())

	bangunDatar = SegitigaSamaSisi{Alas: 6, Tinggi: 5}
	fmt.Printf("Luas Segitiga : %d\n", bangunDatar.Luas())
	fmt.Printf("Keliling Segitiga : %d\n", bangunDatar.Keliling())

	var bangunRuang HitungBangunRuang = Tabung{JariJari: 7, Tinggi: 18}
	fmt.Printf("Luas Permukaan Tabung : %.2f\n", bangunRuang.LuasPermukaan())
	fmt.Printf("Volume Tabung : %.2f\n", bangunRuang.Volume())

	bangunRuang = Balok{Panjang: 28, Lebar: 24, Tinggi: 10}
	fmt.Printf("Luas Permukaan Balok : %.2f\n", bangunRuang.LuasPermukaan())
	fmt.Printf("Volume Balok : %.2f\n", bangunRuang.Volume())

	fmt.Println("==================")

	// soal 2
	var ph PhoneInterface = Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Printf("%s", ph.Display())

	fmt.Println("==================")

	// soal 3
	fmt.Println(LuasPersegi(4, true))

	fmt.Println(LuasPersegi(8, false))

	fmt.Println(LuasPersegi(0, true))

	fmt.Println(LuasPersegi(0, false))

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
