package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type SkinControllerImpl struct {
	Repo repositories.SkinRepository
}

func (s SkinControllerImpl) SelectCarSkins(carId uint) ([]models.Skin, error) {
	return s.Repo.SelectCarSkins(carId)
}

func (s SkinControllerImpl) GetAllSkins() ([]models.Skin, error) {
	return s.Repo.GetAllSkins()
}

func (s SkinControllerImpl) AddSkin(skin models.Skin) error {
	return s.Repo.AddSkin(skin)
}

func (s SkinControllerImpl) UpdateSkin(skin models.Skin) error {
	return s.Repo.UpdateSkin(skin)
}
