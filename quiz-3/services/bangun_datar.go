package services

import "math"

// SEGITIGA SAMA SISI
func LuasSegitigaSamaSisi(c *chan float64, a float64, t float64) {
	*c <- float64(0.5) * a * t 
	close(*c)
}

func KelilingSegitigaSamaSisi(c *chan float64, a float64, t float64) {
	*c <- float64(3) * a
	close(*c)
}

// PERSEGI
func LuasPersegi(c *chan float64, s float64) {
	*c <- float64(s) * float64(s)
	close(*c)
}

func KelilingPersegi(c *chan float64, s float64) {
	*c <- float64(4) * s
	close(*c)
}

// PERSEGI PANJANG
func LuasPersegiPanjang(c *chan float64, p float64, l float64) {
	*c <- p * l
	close(*c)
}

func KelilingPersegiPanjang(c *chan float64, p float64, l float64) {
	*c <- float64(2) * p + float64(2) * l
	close(*c)
}

// LINGKARAN
func LuasLingkaran(c *chan float64, r float64) {
	*c <- float64(math.Pi) * math.Pow(r, 2)
	close(*c)
}

func KelilingLingkaran(c *chan float64, r float64) {
	*c <- float64(2) * math.Pi * r
	close(*c)
}

// JAJAR GENJANG
func LuasJajarGenjang(c *chan float64, a float64, t float64) {
	*c <- a * t
	close(*c)
}

func KelilingJajarGenjang(c *chan float64, a float64, t float64, s float64) {
	*c <- float64(2) * (a + s)
	close(*c)
}