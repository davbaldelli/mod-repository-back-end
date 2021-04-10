package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type NationControllerImpl struct {
	Repo repositories.NationRepository
}

func (n NationControllerImpl) GetAllTracksNations() ([]entities.Nation,error) {
	return n.Repo.SelectAllTrackNations()
}

func (n NationControllerImpl) GetAllBrandsNations() ([]entities.Nation,error) {
	return n.Repo.SelectAllBrandsNations()
}
