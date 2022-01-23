package controllers

import (
	"net/http"
	"quiz-3/services"
	"quiz-3/utils"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func SegitigaSamaSisi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Segitiga Sama Sisi\n"))

	query := r.URL.Query()

	a, _ := strconv.ParseFloat(query.Get("alas"), 64)
	t, _ := strconv.ParseFloat(query.Get("tinggi"), 64)
	c := make(chan float64)

	if query.Get("hitung") == "luas" {
		go services.LuasSegitigaSamaSisi(&c, a, t)
	} else if query.Get("hitung") == "keliling" {
		go services.KelilingSegitigaSamaSisi(&c, a, t)
	}

	utils.ResponseJSON(w, <-c, http.StatusOK)
}

func Persegi(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Persegi\n"))

	query := r.URL.Query()

	s, _ := strconv.ParseFloat(query.Get("sisi"), 64)
	c := make(chan float64)

	if query.Get("hitung") == "luas" {
		go services.LuasPersegi(&c, s)
	} else if query.Get("hitung") == "keliling" {
		go services.KelilingPersegi(&c, s)
	}

	utils.ResponseJSON(w, <-c, http.StatusOK)
}

func PersegiPanjang(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Persegi Panjang\n"))

	query := r.URL.Query()

	panjang, _ := strconv.ParseFloat(query.Get("panjang"), 64)
	l, _ := strconv.ParseFloat(query.Get("lebar"), 64)
	c := make(chan float64)

	if query.Get("hitung") == "luas" {
		go services.LuasPersegiPanjang(&c, panjang, l)
	} else if query.Get("hitung") == "keliling" {
		go services.KelilingPersegiPanjang(&c, panjang, l)
	}

	utils.ResponseJSON(w, <-c, http.StatusOK)
}

func Lingkaran(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Lingkaran\n"))

	query := r.URL.Query()

	rad, _ := strconv.ParseFloat(query.Get("jariJari"), 64)
	c := make(chan float64)

	if query.Get("hitung") == "luas" {
		go services.LuasLingkaran(&c, rad)
	} else if query.Get("hitung") == "keliling" {
		go services.KelilingLingkaran(&c, rad)
	}

	utils.ResponseJSON(w, <-c, http.StatusOK)
}

func JajarGenjang(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("Jajar Genjang\n"))

	query := r.URL.Query()

	alas, _ := strconv.ParseFloat(query.Get("alas"), 64)
	sisi, _ := strconv.ParseFloat(query.Get("sisi"), 64)
	tinggi, _ := strconv.ParseFloat(query.Get("tinggi"), 64)
	c := make(chan float64)

	if query.Get("hitung") == "luas" {
		go services.LuasJajarGenjang(&c, alas, tinggi)
	} else if query.Get("hitung") == "keliling" {
		go services.KelilingJajarGenjang(&c, alas, tinggi, sisi)
	}

	utils.ResponseJSON(w, <-c, http.StatusOK)
}
