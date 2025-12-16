package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is not found in the environment")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to databases:", err)
	}

	apiConfig := apiConfig{
		DB: database.New(conn),
	}

	router := chi.NewRouter()

	// Cors configurations
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,

	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerHealthz)
	v1Router.Get("/error", handlerError)
	// Use routes
	v1Router.Post("/users", apiConfig.handlerCreateUser)
	v1Router.Get("/users/me", apiConfig.midddlewareAuth(apiConfig.handlerGetUser))
	v1Router.Post("/feeds", apiConfig.midddlewareAuth(apiConfig.handlerCreateFeed))
	v1Router.Get("/feeds", apiConfig.handlerGetFeeds)
	v1Router.Get("/feeds/me", apiConfig.midddlewareAuth(apiConfig.handlerGetUserFeeds))
	v1Router.Post("/feeds_follows", apiConfig.midddlewareAuth(apiConfig.handlerCreateFeedFollow))

	// Mount base router to v1 router
	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server starting on Port: %v", port)

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}