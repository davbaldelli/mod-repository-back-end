package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type NationControllerImpl struct {
	Repo repositories.NationRepository
}

func (n NationControllerImpl) GetAllTracksNations() []entities.Nation {
	return n.Repo.GetAllTrackNations()
}

func (n NationControllerImpl) GetAllBrandsNations() []entities.Nation {
	return n.Repo.GetAllBrandsNations()
}
