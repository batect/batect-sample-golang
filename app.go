package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", handlePing)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handlePing(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "pong")
}
