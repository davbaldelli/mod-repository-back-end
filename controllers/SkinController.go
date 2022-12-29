package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type SkinControllerImpl struct {
	Repo repositories.SkinRepository
}

func (s SkinControllerImpl) SelectCarSkins(carId uint) ([]entities.Skin, error) {
	return s.Repo.SelectCarSkins(carId)
}

func (s SkinControllerImpl) AddSkin(skin entities.Skin) error {
	return s.Repo.AddSkin(skin)
}

func (s SkinControllerImpl) UpdateSkin(skin entities.Skin) error {
	return s.Repo.UpdateSkin(skin)
}
