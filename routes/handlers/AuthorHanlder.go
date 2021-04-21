package handlers

import (
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type AuthorHandlerImpl struct {
	AuthorsCtrl controllers.AuthorController
}

func (a AuthorHandlerImpl) GETAllAuthors(writer http.ResponseWriter, request *http.Request) {
	if authors, err := a.AuthorsCtrl.GetAllAuthors(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, authors)
	}
}

func (a AuthorHandlerImpl) GETTrackAuthors(writer http.ResponseWriter, request *http.Request) {
	if authors, err := a.AuthorsCtrl.GetAllTrackAuthors(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, authors)
	}
}

func (a AuthorHandlerImpl) GETCarAuthors(writer http.ResponseWriter, request *http.Request) {
	if authors, err := a.AuthorsCtrl.GetAllCarAuthors(); err != nil {
		respondError(writer, http.StatusInternalServerError, err)
	} else {
		respondJSON(writer, http.StatusOK, authors)
	}
}

