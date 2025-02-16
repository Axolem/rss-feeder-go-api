package main

import "net/http"

func handleHealth(w http.ResponseWriter, _ *http.Request) {
  respondWithJSON(w, 200,
    struct {
    }{})
}
