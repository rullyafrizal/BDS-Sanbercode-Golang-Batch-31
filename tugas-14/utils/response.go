package utils

import (
	"encoding/json"
	"net/http"
)

func ResponseJSON(w http.ResponseWriter, p interface{}, status int) {
	byted, err := json.Marshal(p)
	w.Header().Set("Content-Type", "application/json")
  
	if err != nil {
	  http.Error(w, "Error : ", http.StatusBadRequest)
	}
	
	w.WriteHeader(status)
	w.Write([]byte(byted))
  }