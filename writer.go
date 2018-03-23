package potion

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

// WriteJSON writes JSON struct to response.
func WriteJSON(w http.ResponseWriter, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		glog.Errorf("Failed to encode JSON: %s", err)
		Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
