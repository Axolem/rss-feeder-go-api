package main

import (
  "encoding/json"
  "fmt"
  "github.com/axolem/rss-feeder/internal/database"
  "github.com/google/uuid"
  "net/http"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
  type parammeters struct {
    Name string `json:"name"`
  }

  decoder := json.NewDecoder(r.Body)

  params := parammeters{}
  err := decoder.Decode(&params)

  if err != nil {
    respondWithError(w, 400, fmt.Sprint("Failed reading the body: ", err))
    return
  }

  user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
    ID:   uuid.New(),
    Name: params.Name,
  })

  if err != nil {
    respondWithError(w, 500, fmt.Sprint("Error creating new user: ", err))
    return
  }

  respondWithJSON(w, 201, user)
}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
  respondWithJSON(w, 201, user)
}
