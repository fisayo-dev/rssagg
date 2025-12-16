package main

import (
	"fmt"
	"net/http"

	"github.com/fisayo-dev/rssagg/auth"
	"github.com/fisayo-dev/rssagg/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) midddlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(),apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Unable to get user with API KEY: %v", err))
			return
		}

		handler(w, r, user)
	}
}