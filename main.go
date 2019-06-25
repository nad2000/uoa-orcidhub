//go:generate swagger generate client -f ./assets/uoa-employment-v1.json
//go:generate swagger generate client -f ./assets/orcidhub-api-v0.1-spec.json -A orcidhub
//go:generate swagger generate server -A uoa-orcidhub_

package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
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

const schema = `
CREATE TABLE IF NOT EXISTS user (
	user_id    INTEGER PRIMARY KEY,
    first_name VARCHAR(80)  DEFAULT '',
    last_name  VARCHAR(80)  DEFAULT '',
	email      VARCHAR(250) DEFAULT '',
	password   VARCHAR(250) DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS log (
	id    INTEGER PRIMARY KEY,
	ts	TIMESTAMP DEFAULT current_timestamp, 
	method VARCHAR(10),
	url VARCHAR(1000),
	message TEXT NULL DEFAULT NULL
);
`

type User struct {
	UserId    int    `db:"user_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
	Password  sql.NullString
}

func main() {

	// TODO: figure out how to prefix log lines with a prefix
	logger := log.New()
	logger.Out = os.Stderr
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	logger.Infoln("Starting the server...")

	db, err := sqlx.Connect("sqlite3", "data.db")
	if err != nil {
		log.Fatalln(err)
	}
	db.MustExec(schema)

	h := homepage.NewHandlers(logger, db)
	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	logger.Infof("Certificate: %q", CertFile)
	logger.Infof("Server Key: %q", KeyFile)
	logger.Infof("Service Address: %q", ServiceAddr)
	// err := http.ListenAndServe(":8080", mux) // Invokes the built-in server

	srv := server.New(mux, ServiceAddr)
	err = srv.ListenAndServeTLS(CertFile, KeyFile)
	if err != nil {
		log.WithError(err).Errorln("Failed to start the server.")
	}
}
