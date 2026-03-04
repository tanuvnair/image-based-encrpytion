package handlers

import (
	"encoding/hex"
	"fmt"
	"net/http"
	"strconv"
)

// RandomSource is anything that can produce n random bytes.
type RandomSource interface {
	RandomBytes(n int) ([]byte, error)
}

// HandleRandom returns an HTTP handler that serves random bytes generated
// by the RandomSource. The number of bytes can be controlled via the `bytes`
// query parameter (default: 32, max: 1024).
func HandleRandom(src RandomSource) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := 32 // default byte count

		if raw := r.URL.Query().Get("bytes"); raw != "" {
			parsed, err := strconv.Atoi(raw)
			if err != nil || parsed <= 0 {
				http.Error(w, "invalid `bytes` parameter", http.StatusBadRequest)
				return
			}
			if parsed > 1024 {
				http.Error(w, "`bytes` cannot exceed 1024", http.StatusBadRequest)
				return
			}
			n = parsed
		}

		buf, err := src.RandomBytes(n)
		if err != nil {
			http.Error(w, "failed to generate random bytes", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "random bytes (%d): %s\n", n, hex.EncodeToString(buf))
	}
}
