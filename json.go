package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Responding with 5xx error: ", message)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, 500, errResponse{
		Error: message,
	})
}
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal JSON response: %v\n", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	_, wErr := w.Write(data)
	if wErr != nil {
		log.Println("Failed to respond")
		return
	}
}
