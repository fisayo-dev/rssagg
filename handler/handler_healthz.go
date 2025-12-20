package handler

import (
	"net/http"

	"github.com/fisayo-dev/rssagg/database"
	"github.com/fisayo-dev/rssagg/utils"
)

type ApiConfig struct {
	DB *database.Queries
	
}
func HandlerHealthz(w http.ResponseWriter, r *http.Request){
	utils.RespondWithJSON(w, 200, struct{}{})
}