package homepage

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

const message = "..."

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	// microoptimization:
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// middlewhare
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next(w, r)
		h.logger.Infof("Request processed in %d", time.Now().Sub(startTime))
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {

	// simple middleware implementation (it can be done with Chi or Gorilla too)
	// but if you want to run this thing as a lamba w/o other external dependencies
	// this would be sufficient
	mux.HandleFunc("/", h.Logger(h.Home))
	mux.HandleFunc("/home", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
