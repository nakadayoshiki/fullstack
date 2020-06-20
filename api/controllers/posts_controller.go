package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/auth"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/models"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/responses"
	"github.com/nakadayoshiki/fullstack/github.com/nakadayoshiki/fullstack/api/utils/formaterror"
)

func (s *Server) CreatedPost(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	uid, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}
	if uid != post.AuthorID {
		responses.ERROR(w, http.StatusUnauthorized, errors.New(http.StatusText(http.StatusUnauthorized)))
		return
	}

	postCreated, err := post.SavePost(s.DB)
	if err != nil {
		formatedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formatedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.URL.Path, postCreated.ID))
	responses.JSON(w, http.StatusCreated, postCreated)
}

func (s *Server) GetPosts(w http.RsponseWriter, r *http.Request) {
	post := models.Post{}
	posts, err := post.FindAllPosts(s.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, posts)
}
