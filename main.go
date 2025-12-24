package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	// "time"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/fisayo-dev/rssagg/handler"
	// "github.com/fisayo-dev/rssagg/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)


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

	apiCfg := handler.ApiConfig{
		DB: database.New(conn),
	}

	// Call scraping func
	// go utils.StartScraping(apiCfg.DB,10,time.Minute)

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
	v1Router.Get("/healthz", handler.HandlerHealthz)
	v1Router.Get("/error", handler.HandlerError)
	// Use routes
	v1Router.Post("/users", apiCfg.HandlerCreateUser)
	v1Router.Get("/users/me", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUser))
	v1Router.Post("/feeds", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.HandlerGetFeeds)
	v1Router.Get("/feeds/me", apiCfg.MiddlewareAuth(apiCfg.HandlerGetUserFeeds))
	v1Router.Post("/feeds_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerCreateFeedFollow))
	v1Router.Get("/feeds_follows", apiCfg.MiddlewareAuth(apiCfg.HandlerGetFeedFollow))
	v1Router.Delete("/feeds_follows/{feedFollowID}", apiCfg.MiddlewareAuth(apiCfg.HandlerDeleteFeedFollow))
	v1Router.Get("/posts", apiCfg.MiddlewareAuth(apiCfg.HandlerGetPostsForUser))

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