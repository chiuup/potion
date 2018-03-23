package potion

import (
	"encoding/json"
	"net/http"
)

type jsonError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}

// Error writes an formatted error
func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(jsonError{code, error})
}
