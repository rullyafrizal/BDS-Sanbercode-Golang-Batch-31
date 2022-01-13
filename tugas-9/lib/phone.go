package lib

import (
	"strconv"
	"strings"
)

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (ph Phone) Display() string {
	return "name : " + ph.Name + "\n" + "brand : " + ph.Brand + "\n" + "year : " + strconv.Itoa(ph.Year) + "\n" + "colors : " + strings.Join(ph.Colors, ", ") + "\n"
}