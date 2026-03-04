package handlers

import (
	"fmt"
	"net/http"
)

func HandleRandom(w http.ResponseWriter, r *http.Request) {
	// TODO: generate random bytes from the CSPRNG
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintln(w, "random bytes will be served here")
}
