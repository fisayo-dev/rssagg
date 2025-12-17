package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/fisayo-dev/rssagg/utils"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

// Made handlerCreateUser() a method because we want to pass in apiCfg
// and we can't change the structure of or handle function in go
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	// Create an instance of parameters struct
	params := parameters{}
	// Create new JSON decoder
	decoder := json.NewDecoder(r.Body)
	// Parse the JSON request body into the parameters struct
	err := decoder.Decode(&params)
	if err != nil {
		utils.RespondWithError(w,400,fmt.Sprintf("Error parsing json: %v", err))
		return 
	}
	
	follow_feed, err := apiCfg.DB.FollowFeed(r.Context(), database.FollowFeedParams{
		ID: uuid.New(),
		UserID: user.ID,
		FeedID: params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Unable to follow feed, id: %v error: %v", params.FeedID, err))
		return
	}
	utils.RespondWithJSON(w, 201, databaseFeedFollowToFeedFollow(follow_feed))
}

func (apiCfg *apiConfig) handlerGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Unable to get user fess followed: %v", err))
		return
	}
	utils.RespondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (apiCfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User){
	feedFollowIDstr := chi.URLParam(r, "feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDstr)
	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Unable to parse feed follow id: %v", err))
		return
	}

	// Delete follow feed
	err = apiCfg.DB.DeleteFeedFollow(r.Context(),database.DeleteFeedFollowParams{
		ID: feedFollowID,
		UserID: user.ID,
	})

	if err != nil {
		utils.RespondWithError(w, 400, fmt.Sprintf("Unable to delete feed follow: %v", err))
		return;
	}

	utils.RespondWithJSON(w, 200, struct{}{})
}
