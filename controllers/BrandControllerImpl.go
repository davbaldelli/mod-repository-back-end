package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	repo repositories.BrandRepository
}

func (b BrandControllerImpl) GetAllBrands() []models.CarBrand {
	return b.repo.GetAllBrands()
}

func (b BrandControllerImpl) GetBrandByNation(nationName string) []models.CarBrand {
	return b.repo.GetBrandByNation(nationName)
}

func (b BrandControllerImpl) GetBrandByName(name string) []models.CarBrand {
	return b.repo.GetBrandByName(name)
}

func (b BrandControllerImpl) AddBrand(name string, nation models.Nation) error {
	return b.repo.AddNewBrand(models.CarBrand{
		Name:   name,
		Nation: nation,
	})
}
