package homepage

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

const message = "..."

type Handlers struct {
	logger *log.Logger
	db     *sqlx.DB
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
		message := fmt.Sprintf("Request processed in %d", time.Now().Sub(startTime))
		h.logger.Info(message)
		// use the context to provide ability to cancel the handing and query too:
		_, err := h.db.ExecContext(r.Context(), "INSERT INTO log (method, url, message) VALUES ($1, $2, $3)",
			r.Method, r.RequestURI, message)
		if err != nil {
			h.logger.WithError(err).Error("Failed to insert a record.")
		}
	}
}

func (h *Handlers) SetupRoutes(mux *http.ServeMux) {

	// simple middleware implementation (it can be done with Chi or Gorilla too)
	// but if you want to run this thing as a lamba w/o other external dependencies
	// this would be sufficient
	mux.HandleFunc("/", h.Logger(h.Home))
	mux.HandleFunc("/home", h.Logger(h.Home))
}

func NewHandlers(logger *log.Logger, db *sqlx.DB) *Handlers {
	return &Handlers{
		logger: logger,
		db:     db,
	}
}
