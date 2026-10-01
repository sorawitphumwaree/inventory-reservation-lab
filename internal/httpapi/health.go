package httpapi

import (
	"encoding/json/v2"
	"log"
	"net/http"
)

type healthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := healthResponse{
		Status: "ok",
	}

	w.Header().Set("Content-Type", "application/json")

	err := json.MarshalWrite(w, response)
	if err != nil {
		log.Printf("write health response: %v", err)
	}
}
