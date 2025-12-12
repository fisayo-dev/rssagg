package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT is not found in the environment")
	}
	fmt.Printf("PORT running on: %v", port)
	
}