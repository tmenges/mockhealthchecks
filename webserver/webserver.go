package main

import (
	"fmt"
	"log"
	"net/http"
)

type readinessHandler struct {
}

type livenessHandler struct {
}

func (h *readinessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Readiness Handler called")
	fmt.Fprintf(w, "ready\n")
}

func (h *livenessHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Print("Liveness Handler called")
	fmt.Fprintf(w, "alive\n")
}

func main() {
	http.Handle("/readiness", new(readinessHandler))
	http.Handle("/liveness", new(livenessHandler))
	log.Fatal(http.ListenAndServe(":8338", nil))
}
