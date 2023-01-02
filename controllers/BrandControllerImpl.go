package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	Repo repositories.BrandRepository
}

func (b BrandControllerImpl) GetAllBrands() ([]models.CarBrand, error) {
	return b.Repo.SelectAllBrands()
}
