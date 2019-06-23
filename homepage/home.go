package homepage

import "net/http"

const message = "..."

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// microoptimization:
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
