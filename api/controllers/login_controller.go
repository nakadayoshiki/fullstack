package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/auth"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/models"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/responses"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/utils/formaterror"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	token, err := s.SignIn(user.Email, user.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (s *Server) SignIn(email, password string) (string, error) {
	var err error
	user := models.User{}
	err = s.DB.Debug().Model(models.User{}).Where("id = ?", email).Take(&user).Error
	if err != nil {
		return "", err
	}

	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}
