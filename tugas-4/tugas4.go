package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding\n", i)
			} else {
				fmt.Printf("%d - Santai\n", i)
			}
		} else {
			fmt.Printf("%d - Berkualitas\n", i)
		}
	}

	dash := strings.Repeat("=", 25)
	fmt.Printf("\n%s\n", dash)

	// soal 2
	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}

		fmt.Printf("\n")
	}

	fmt.Printf("\n%s\n", dash)

	// soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	for i, item := range kalimat {
		if i > 1 {
			fmt.Printf("%s ", item)
		}
	}

	fmt.Println("")
	fmt.Printf("\n%s\n", dash)

	// soal 4
	var sayuran = []string{}

	sayuran = append(
		sayuran,
		"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun",
	)

	for i, item := range sayuran {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	fmt.Printf("\n%s\n", dash)

	// soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	satuan["volume balok"] = satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]

	for k, v := range satuan {
		fmt.Printf("%s = %d\n", strings.Title(strings.ToLower(k)), v)
	}
}
