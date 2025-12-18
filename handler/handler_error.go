package handler

import (
	"net/http"

	"github.com/fisayo-dev/rssagg/utils"
)

func HandlerError(w http.ResponseWriter, r *http.Request){
	utils.RespondWithError(w, 400, "This is a server error ")
}