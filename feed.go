package main

import (
  "encoding/json"
  "fmt"
  "github.com/axolem/rss-feeder/internal/database"
  "github.com/google/uuid"
  "net/http"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
  type parammeters struct {
    Name string `json:"name"`
    Url  string `json:"url"`
  }

  decoder := json.NewDecoder(r.Body)

  params := parammeters{}
  err := decoder.Decode(&params)

  if err != nil {
    respondWithError(w, 400, fmt.Sprint("Failed reading the body: ", err))
    return
  }

  feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
    ID:     uuid.New(),
    Name:   params.Name,
    Url:    params.Url,
    UserID: user.ID,
  })

  if err != nil {
    respondWithError(w, 500, fmt.Sprint("Error creating new feed: ", err))
    return
  }

  respondWithJSON(w, 201, feed)
}
