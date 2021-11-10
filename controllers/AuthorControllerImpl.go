package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type AuthorsControllerImpl struct {
	Repo repositories.AuthorRepository
}

func (a AuthorsControllerImpl) GetAllAuthors() ([]entities.Author, error) {
	return a.Repo.SelectAllAuthors()
}

func (a AuthorsControllerImpl) GetAllCarAuthors() ([]entities.Author, error) {
	return a.Repo.SelectAllCarAuthors()
}

func (a AuthorsControllerImpl) GetAllTrackAuthors() ([]entities.Author, error) {
	return a.Repo.SelectAllTrackAuthors()
}
