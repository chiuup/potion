package potion

import (
	"net/http"
	"strings"
)

// ContentTypeHandler checks for ContentType and rejects invalid requests.
func ContentTypeHandler(next http.HandlerFunc, contentTypes ...string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for _, contentType := range contentTypes {
			ct := r.Header.Get("Content-Type")
			if i := strings.IndexRune(ct, ';'); i != -1 {
				ct = ct[0:i]
			}
			if ct == contentType {
				next.ServeHTTP(w, r)
				return
			}
		}

		Error(w, http.StatusText(http.StatusUnsupportedMediaType), http.StatusUnsupportedMediaType)
	}
}
