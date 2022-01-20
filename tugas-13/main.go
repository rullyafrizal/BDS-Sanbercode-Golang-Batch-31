package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func main() {
	handle()
}

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json:"id"`
}

type NilaiMahasiswaRequest struct {
	Nama       string `json:"nama"`
	MataKuliah string `json:"mata_kuliah"`
	Nilai      uint   `json:"nilai"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func indeksNilai(n *NilaiMahasiswa) {
	if n.Nilai >= 80 {
		n.IndeksNilai = "A"
	} else if n.Nilai >= 70 && n.Nilai < 80 {
		n.IndeksNilai = "B"
	} else if n.Nilai >= 60 && n.Nilai < 70 {
		n.IndeksNilai = "C"
	} else if n.Nilai >= 50 && n.Nilai < 60 {
		n.IndeksNilai = "D"
	} else if n.Nilai < 50 {
		n.IndeksNilai = "E"
	}
}

func determineId(n *NilaiMahasiswa) {
	if len(nilaiNilaiMahasiswa) == 0 {
		n.ID = 1
	} else {
		n.ID = nilaiNilaiMahasiswa[len(nilaiNilaiMahasiswa)-1].ID + 1
	}
}

func store(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var nilaiMahasiswa NilaiMahasiswa
	var request NilaiMahasiswaRequest

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJson := json.NewDecoder(r.Body)

			if err := decodeJson.Decode(&request); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if request.Nilai > 100 || request.Nilai < 0 {
				http.Error(w, "nilai tidak boleh lebih dari 100 atau lebih kecil dari 0", http.StatusBadRequest)
				return
			}

			nilaiMahasiswa = NilaiMahasiswa{
				Nama:       request.Nama,
				MataKuliah: request.MataKuliah,
				Nilai:      request.Nilai,
			}

			indeksNilai(&nilaiMahasiswa)
			determineId(&nilaiMahasiswa)
		} else {
			nama := r.FormValue("nama")
			mataKuliah := r.FormValue("mataKuliah")
			nilai := r.FormValue("nilai")

			if nilai > "100" || nilai < "0" {
				http.Error(w, "nilai tidak boleh lebih dari 100 atau lebih kecil dari 0", http.StatusBadRequest)
				return
			}

			nilaiMahasiswa.Nama = nama
			nilaiMahasiswa.MataKuliah = mataKuliah
			strNilai, _ := strconv.Atoi(nilai)
			nilaiMahasiswa.Nilai = uint(strNilai)

			indeksNilai(&nilaiMahasiswa)
			determineId(&nilaiMahasiswa)
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilaiMahasiswa)
		dataNilaiMahasiswa, err := json.Marshal(nilaiMahasiswa)

		if err != nil {
			log.Fatal(err)
		}
		w.Write(dataNilaiMahasiswa)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilaiMahasiswa, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilaiMahasiswa)
		return
	}

	http.Error(w, "ERROR....", http.StatusNotFound)
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uname, pwd, ok := r.BasicAuth()
		if !ok {
			w.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if uname == "admin" && pwd == "admin" {
			next.ServeHTTP(w, r)
			return
		}
		w.Write([]byte("Username atau Password tidak sesuai"))
		return
	})
}

func handle() {
	http.Handle("/store-nilai-mahasiswa", Auth(http.HandlerFunc(store)))
	http.HandleFunc("/nilai-mahasiswa", index)

	log.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
