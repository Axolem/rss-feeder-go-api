package main

import "net/http"

func handleRediness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200,
		struct {
		}{})
}
