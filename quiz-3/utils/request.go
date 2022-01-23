package utils

import (
	"encoding/json"
	"net/http"
)

func ParseBody(r *http.Request, v interface{}) error {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(v)

	if err != nil {
		return err
	}

	return nil
}