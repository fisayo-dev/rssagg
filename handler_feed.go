package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/google/uuid"
)

// Made handlerCreateUser() a method because we want to pass in apiCfg
// and we can't change the structure of or handle function in go
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	// Create an instance of parameters struct
	params := parameters{}
	// Create new JSON decoder
	decoder := json.NewDecoder(r.Body)
	// Parse the JSON request body into the parameters struct
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w,400,fmt.Sprintf("Error parsing json: %v", err))
		return 
	}
	
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID: uuid.New(),
		UserID: user.ID,
		Url: params.Url,
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to create feed: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedToFeed(feed))
}

func (apiCfg *apiConfig) handlerGetUserFeeds(w http.ResponseWriter, r *http.Request, user database.User){
	feeds, err := apiCfg.DB.GetUserFeeds(r.Context(),user.ID)
	if err != nil {
		respondWithError(w, 404, fmt.Sprintln("Error occurred finding feeds"))
		return;
	}

	respondWithJSON(w, 200, feeds)
}