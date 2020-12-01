package controllers

import (
	"net/http"

	"github.com/triano4/uploaddrive/api/responses"
)

//Home function
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")

}
