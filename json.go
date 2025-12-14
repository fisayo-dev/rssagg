package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// Convert payload to byte (binary)
	data, err := json.Marshal(payload)
	//  If eror occurs log error
	 if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	 }

	//  Else respond with the header and data
	 w.Header().Add("Content-Type", "application/json")
	 w.WriteHeader(code)
	 w.Write(data)
}