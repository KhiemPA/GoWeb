package utils

import (
	"net/http"
	"encoding/json"
)

func ToJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(data)
}