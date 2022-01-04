package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// soal 1
	var bootcamp string = "Bootcamp"
	var digital string = "Digital"
	var skill string = "Skill"
	var sanbercode string = "Sanbercode"
	var golang string = "Golang"

	fmt.Printf("%s %s %s %s %s\n", bootcamp, digital, skill, sanbercode, golang)

	// soal 2
	halo := "Halo Dunia"
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", 1))

	// soal 3
	var kataPertama string = "saya"
	var kataKedua string = "senang"
	var kataKetiga string = "belajar"
	var kataKeempat string = "golang"

	fmt.Printf("%s %s %s %s\n", kataPertama, strings.Replace(kataKedua, "s", "S", 1), strings.Replace(kataKetiga, "r", "R", 1), strings.ToUpper(kataKeempat))

	// soal 4
	var angkaPertama string = "8"
	var angkaKedua string = "5"
	var angkaKetiga string = "6"
	var angkaKeempat string = "7"

	angkaPert, _ := strconv.ParseInt(angkaPertama, 10, 64)
	angkaKed, _ := strconv.ParseInt(angkaKedua, 10, 64)
	angkaKet, _ := strconv.ParseInt(angkaKetiga, 10, 64)
	angkaEmp, _ := strconv.ParseInt(angkaKeempat, 10, 64)

	fmt.Printf("%d + %d + %d + %d = %d\n", angkaPert, angkaKed, angkaKet, angkaEmp, angkaPert+angkaKed+angkaKet+angkaEmp)

	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	fmt.Printf("\"%s\" - %d\n", strings.Replace(kalimat, "halo", "hi", 2), angka)
}
