package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/fisayo-dev/rssagg/utils"
	"github.com/google/uuid"
)

// Made handlerCreateUser() a method because we want to pass in apiCfg
// and we can't change the structure of or handle function in go
func (apiCfg *ApiConfig) HandlerCreateUser(w http.ResponseWriter, r *http.Request){
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
		utils.RespondWithError(w,400,fmt.Sprintf("Error parsing json: %v", err))
		return 
	}
	hashedPassword, err := utils.HashPassword(params.Password)
	if err != nil {
		utils.RespondWithError(w, 500, fmt.Sprintf("Error hashing password: %v", err))
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
		utils.RespondWithError(w, 400, fmt.Sprintf("Unable to create user: %v", err))
		return
	}
	utils.RespondWithJSON(w, 201, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetUser(w http.ResponseWriter, r *http.Request, user database.User){
	utils.RespondWithJSON(w, 200, utils.DatabaseUserToUser(user))
}

func (apiCfg *ApiConfig) HandlerGetPostsForUser(w http.ResponseWriter, r*http.Request, user database.User){
	posts, err := apiCfg.DB.GetPostsForUser(r.Context(),database.GetPostsForUserParams{
		UserID: user.ID,
		Limit: 10,
	})
	if err != nil {
		utils.RespondWithError(w,400,fmt.Sprintf("Failed to get posts for user: %v", err))
		return
	}
	utils.RespondWithJSON(w,200,utils.DatabasePostsToPosts(posts))
}