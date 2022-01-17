package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"sync"
	"time"
)

func main() {
	// soal 1
	var phones []string = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i, v := range phones {
			fmt.Printf("%d. %s\n", i+1, v)
			time.Sleep(1 * time.Second)
		}

		wg.Done()
	}()

	wg.Wait()

	// soal 2
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(&moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// soal 3
	kelilingLing := make(chan float64)
	luasLing := make(chan float64)
	volumeTab := make(chan float64)

	var r1 float64 = 8
	var r2 float64 = 14
	var r3 float64 = 20
	var t float64 = 10

	go luasLingkaran(&luasLing, r1, r2, r3)
	go kelilingLingkaran(&kelilingLing, r1, r2, r3)

	go volumeTabung(&volumeTab, t, r1, r2, r3)

	for val := range kelilingLing {
		fmt.Printf("%.2f\n", val)
	}

	for val := range luasLing {
		fmt.Printf("%.2f\n", val)
	}

	for val := range volumeTab {
		fmt.Printf("%.2f\n", val)
	}

	// soal 4
	var p int = 15
	var l int = 12
	var tinggi int = 8

	var ch1 = make(chan int)
	go luasPersegiPanjang(&ch1, p, l)

	var ch2 = make(chan int)
	go kelilingPersegiPanjang(&ch2, p, l)

	var ch3 = make(chan int)
	go volumeBalok(&ch3, p, l, tinggi)

	for i := 0; i < 3; i++ {
		select {
		case luas := <-ch1:
			fmt.Printf("Luas Persegi Panjang : %d \n", luas)
		case keliling := <-ch2:
			fmt.Printf("Keliling Persegi Panjang : %d \n", keliling)
		case volume := <-ch3:
			fmt.Printf("Volume Balok : %d \n", volume)
		}
	}
}

func getMovies(c *chan string, movies ...string) {
	*c <- "List Movies:"
	for i, movie := range movies {
		*c <- strconv.Itoa(i+1) + ". " + movie
	}

	close(*c)
}

func luasLingkaran(c *chan float64, r ...float64) {
	for _, v := range r {
		*c <- float64(math.Pi) * math.Pow(v, 2)
	}
	close(*c)
}

func kelilingLingkaran(c *chan float64, r ...float64) {
	for _, v := range r {
		*c <- float64(math.Pi) * 2 * v
	}
	close(*c)
}

func volumeTabung(c *chan float64, t float64, r ...float64) {
	for _, v := range r {
		*c <- float64(math.Pi) * math.Pow(v, 2) * t
	}

	close(*c)
}

func luasPersegiPanjang(ch *chan int, p int, l int) {
	*ch <- p * l
}

func kelilingPersegiPanjang(ch *chan int, p int, l int) {
	*ch <- 2 * (p + l)
}

func volumeBalok(ch *chan int, p int, l int, t int) {
	*ch <- p * l * t
}
