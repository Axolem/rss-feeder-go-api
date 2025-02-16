package main

import (
	"fmt"
	"github.com/axolem/rss-feeder/internal/auth"
	"github.com/axolem/rss-feeder/internal/database"
	"net/http"
)

type authHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiKey, err := auth.GetApiKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Could not find the user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
