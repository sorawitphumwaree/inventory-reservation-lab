package main

import (
	"github.com/sorawitphumwaree/inventory-reservation-lab/internal/httpapi"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", httpapi.HealthHandler)

	log.Println("server listening on http://127.0.0.1:8080")

	err := http.ListenAndServe("127.0.0.1:8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
