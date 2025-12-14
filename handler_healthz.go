package main

import "net/http"

func handlerHealthz(w http.ResponseWriter, r *http.Request){
	respondWithJSON(w, 200, struct{}{})
}