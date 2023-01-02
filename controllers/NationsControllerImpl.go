package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type NationControllerImpl struct {
	Repo repositories.NationRepository
}

func (n NationControllerImpl) GetAllTracksNations() ([]models.Nation, error) {
	return n.Repo.SelectAllTrackNations()
}

func (n NationControllerImpl) GetAllBrandsNations() ([]models.Nation, error) {
	return n.Repo.SelectAllBrandsNations()
}
