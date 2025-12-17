package main

import (
	"net/http"

	"github.com/fisayo-dev/rssagg/utils"
)

func handlerHealthz(w http.ResponseWriter, r *http.Request){
	utils.RespondWithJSON(w, 200, struct{}{})
}