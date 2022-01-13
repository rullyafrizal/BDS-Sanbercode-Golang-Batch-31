package lib

import "math"

// Tabung
type Tabung struct {
	JariJari, Tinggi float64
}

func (t Tabung) Volume() float64 {
	return math.Pi * math.Pow(t.JariJari, 2) * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	var luasTutup float64 = 2 * math.Pi * math.Pow(t.JariJari, 2)
	var panjangSelimut float64 = 2 * math.Pi * t.JariJari
	var luasSelimut float64 = panjangSelimut * t.Tinggi

	return luasTutup + luasSelimut
}