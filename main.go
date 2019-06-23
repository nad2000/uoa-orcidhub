//go:generate swagger generate client -f ./assets/uoa-employment-v1.json
//go:generate swagger generate client -f ./assets/orcidhub-api-v0.1-spec.json -A orcidhub
//go:generate swagger generate server -A uoa-orcidhub_

package main

import (
	"net/http"
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/nad2000/uoa-orcidhub/server"
)

const message = "..."

func getenv(key, defaultValue string) {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = defaultValue
	}
	return value
}

var (
	CertFile    = os.Getenv("CERT_FILE")
	KeyFile     = os.Getenv("KEY_FILE")
	ServiceAddr = os.Getenv("SERVICE_ADDR")
)

func main() {

	if CertFile == "" {
		CertFile 
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// microoptimization:
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(message))
	})

	log.Infoln("Starting the server...")
	// err := http.ListenAndServe(":8080", mux) // Invokes the built-in server

	srv := server.New(mux, ServiceAddr)
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.WithError(err).Errorln("Failed to start the server.")
	}
}
