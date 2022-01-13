package lib

// Balok
type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * ((b.Panjang * b.Lebar) + (b.Panjang * b.Tinggi) + (b.Lebar * b.Tinggi)))
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}
