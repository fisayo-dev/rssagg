package main

import "net/http"

func handlerError(w http.ResponseWriter, r *http.Request){
	respondWithError(w, 400, "This is a server error ")
}