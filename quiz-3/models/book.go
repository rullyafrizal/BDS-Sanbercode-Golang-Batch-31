package models

import (
	"net/url"
	"strconv"
	"time"
)

type Book struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url"`
	ReleaseYear int    `json:"release_year"`
	Price       string `json:"price"`
	TotalPage   int    `json:"total_page"`
	Thickness   string `json:"thickness"`
	CategoryId  int    `json:"category_id"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func (b *Book) ValidateBook() map[string]string {
	var errMsg = make(map[string]string)

	if !isValidUrl(b.ImageUrl) {
		errMsg["image_url"] = "Invalid image url (ex: https://picsum.photos/200/300)"
	}

	if b.ReleaseYear < 1980 || b.ReleaseYear > time.Now().Year() {
		errMsg["release_year"] = "Invalid release year, must be between 1980-" + strconv.Itoa(time.Now().Year())
	}

	return errMsg
}

func (b *Book) ConvertBook() {
	if b.TotalPage <= 100 {
		b.Thickness = "tipis"
	} else if b.TotalPage <= 200 && b.TotalPage > 100 {
		b.Thickness = "sedang"
	} else {
		b.Thickness = "tebal"
	}
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
