package controllers

import (
	"net/http"

	"github.com/rian-io/GoBlog/api/responses"
)

// Home ...
func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To This Awesome API")
}
