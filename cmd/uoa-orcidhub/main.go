package main

import (
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Message struct {
	Subject string `json:"subject"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func echo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `{"now": %q}`, time.Now().Format(time.RFC3339))
}

func handleUpdate(w http.ResponseWriter, r *http.Request) {
}

func handleRequests() {
	http.HandleFunc("/echo", echo)
	// http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
