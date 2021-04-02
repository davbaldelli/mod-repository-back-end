package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type NationControllerImpl struct {
	repo repositories.NationRepository
}

func (n NationControllerImpl) GetAllTracksNations() []models.Nation {
	return n.repo.GetAllTrackNations()
}

func (n NationControllerImpl) GetAllBrandsNations() []models.Nation {
	return n.repo.GetAllBrandsNations()
}
