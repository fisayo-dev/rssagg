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
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request){
	type parameters struct {
		Name string `json:"name"`
		Email string `json:"email"`
		Password string `json:"password"`
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
	hashedPassword, err := hashPassword(params.Password)
	if err != nil {
		respondWithError(w, 500, fmt.Sprintf("Error hashing password: %v", err))
		return
	}
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		Email: params.Email,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Password: hashedPassword,

	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetPostsForUser(w http.ResponseWriter, r*http.Request, user database.User){
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		respondWithError(w,40,fmt.Sprintf("Failed to get posts for user: %v", err))
	}
	respondWithJSON(w,200,databasePostsToPosts(posts))
}