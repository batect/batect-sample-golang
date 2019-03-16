package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	address := "0.0.0.0:8080"

	log.Infof("Starting on %s...", address)

	http.HandleFunc("/ping", handlePing)
	log.Fatal(http.ListenAndServe(address, nil))
}

func handlePing(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "pong")
}
