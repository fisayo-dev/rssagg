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