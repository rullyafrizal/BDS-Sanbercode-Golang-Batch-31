package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

func main() {
	// soal 3
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	// soal 1
	defer tampilKalimat("Golang Backend Development", 2021)

	// soal 2
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic: ", r)
		} else {
			fmt.Println("App running perfectly")
		}
	}()

	if s, err := kelilingSegitigaSamakeliling(2, true); err != nil {
		panic(err.Error())
	} else {
		fmt.Println(s)
	}

	// soal 4
	var phones = []string{}

	tambahPhone(&phones, "Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo")

	sort.Strings(phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second * 1)
	}

	// soal 5
	var radius []float64 = []float64{7, 10, 15}

	for _, r := range radius {
		fmt.Println("Luas Lingkaran dengan radius", r, "adalah", luasLingkaran(r))
		fmt.Println("Keliling Lingkaran dengan radius", r, "adalah", kelilingLingkaran(r))
	}

	// soal 6
	var panjang = flag.Int("panjang", 10, "Panjang Persegi Panjang")
	var lebar = flag.Int("lebar", 6, "Lebar Persegi Panjang")

	flag.Parse()
	fmt.Println("Luas Persegi Panjang dengan panjang", *panjang, "dan lebar", *lebar, "adalah", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling Persegi Panjang dengan panjang", *panjang, "dan lebar", *lebar, "adalah", kelilingPersegiPanjang(*panjang, *lebar))
}

func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func luasLingkaran(r float64) float64 {
	return math.Ceil(math.Pi * math.Pow(r, 2))
}

func kelilingLingkaran(r float64) float64 {
	return math.Ceil(math.Pi * 2 * r)
}

func tambahPhone(phones *[]string, phone ...string) {
	*phones = append(*phones, phone...)
}

func cetakAngka(angka *int) {
	fmt.Printf("Total Angka : %d\n", *angka)
}

func tambahAngka(x int, angka *int) {
	*angka += x
}

func tampilKalimat(s string, x int) {
	fmt.Printf("Isi kalimatnya adalah %s\n", s)

	fmt.Printf("Isi tahunnya adalah %d\n", x)
}

func kelilingSegitigaSamakeliling(n int, x bool) (string, error) {
	var keliling int = n * 3

	if keliling != 0 {
		if x {
			return "keliling segitiga sama sisi dengan sisinya " + strconv.Itoa(keliling) + " cm adalah " + strconv.Itoa(keliling) + " cm", nil
		}

		return strconv.Itoa(keliling), nil
	} else {
		if x {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}

	return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
}
