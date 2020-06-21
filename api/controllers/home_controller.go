package controllers

import (
	"net/http"

	"github.com/nakadayoshiki/fullstack/api/responses"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcomu To This Awesome API")
}
