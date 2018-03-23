package potion

import (
	"encoding/json"
	"net/http"

	"github.com/golang/glog"
)

// ReadJSON reads json in body
func ReadJSON(w http.ResponseWriter, r *http.Request, v interface{}) bool {
	if r.Body == nil {
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		glog.Errorf("Failed to read JSON: %s", err)
		Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return false
	}
	return true
}
