package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}

	router := chi.NewRouter()

	// Cors configurations
	// router.Use(cors.Handler(cors.Options{
	// 	AllowedOrigins: []string{"https://*", "http://*"},
	// 	AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	// 	AllowedHeaders: []string{"*"},
	// 	ExposdHeaders: []string{"Link"},
	// 	AllowedCredentials: false,
	// 	MaxAfge: 300,

	// }))
	server := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf(":%s", port),
	}

	log.Printf("Server starting on Port: %v", port)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}