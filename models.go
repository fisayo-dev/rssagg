package main

import (
	"time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/google/uuid"
)


type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	ApiKey string `json:"api_key"`
}
type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url     string `json:"url"`
	UserID  uuid.UUID `json:"user_id"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID: dbUser.ID,
		Email: dbUser.Email,
		Password: dbUser.Password,
		Name: dbUser.Name,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		ApiKey: dbUser.ApiKey,
	}
}
func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID: dbFeed.ID,
		Url: dbFeed.Url,
		Name: dbFeed.Name,
		UserID: dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}