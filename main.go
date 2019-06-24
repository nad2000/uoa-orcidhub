//go:generate swagger generate client -f ./assets/uoa-employment-v1.json
//go:generate swagger generate client -f ./assets/orcidhub-api-v0.1-spec.json -A orcidhub
//go:generate swagger generate server -A uoa-orcidhub_

package main

import (
	"net/http"
	"os"

	"github.com/nad2000/uoa-orcidhub/homepage"
	"github.com/nad2000/uoa-orcidhub/server"
	log "github.com/sirupsen/logrus"
)

func getenv(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		value = defaultValue
	}
	return value
}

var (
	CertFile    = getenv("CERT_FILE", "./server.crt")
	KeyFile     = getenv("KEY_FILE", "./server.key")
	ServiceAddr = getenv("SERVICE_ADDR", ":8080")
)

func main() {

	// TODO: figure out how to prefix log lines with a prefix
	logger := log.New()
	logger.Out = os.Stderr
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	logger.Infoln("Starting the server...")
	logger.Infof("Certificate: %q", CertFile)
	logger.Infof("Server Key: %q", KeyFile)
	logger.Infof("Service Address: %q", ServiceAddr)
	// err := http.ListenAndServe(":8080", mux) // Invokes the built-in server

	srv := server.New(mux, ServiceAddr)
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.WithError(err).Errorln("Failed to start the server.")
	}
}
