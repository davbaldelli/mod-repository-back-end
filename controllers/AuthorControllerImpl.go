package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type AuthorsControllerImpl struct {
	Repo repositories.AuthorRepository
}

func (a AuthorsControllerImpl) GetAllAuthors() ([]models.Author, error) {
	return a.Repo.SelectAllAuthors()
}

func (a AuthorsControllerImpl) GetAllCarAuthors() ([]models.Author, error) {
	return a.Repo.SelectAllCarAuthors()
}

func (a AuthorsControllerImpl) GetAllTrackAuthors() ([]models.Author, error) {
	return a.Repo.SelectAllTrackAuthors()
}
