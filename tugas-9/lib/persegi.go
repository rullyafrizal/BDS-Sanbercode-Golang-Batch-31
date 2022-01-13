package lib

import "strconv"

func LuasPersegi(sisi int, x bool) interface{} {
	var luas int = sisi * sisi

	if sisi != 0 {
		if x {
			return "luas persegi dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(luas) + " cm"
		}

		return sisi
	} else {
		if x {
			return "Maaf anda belum menginput sisi persegi"
		}
	}

	return nil
}