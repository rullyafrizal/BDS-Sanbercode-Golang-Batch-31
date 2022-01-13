package lib

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type PhoneInterface interface {
	Display() string
}
