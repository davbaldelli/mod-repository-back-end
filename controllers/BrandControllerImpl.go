package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	Repo repositories.BrandRepository
}

func (b BrandControllerImpl) GetAllBrands() []models.CarBrand {
	return b.Repo.GetAllBrands()
}

func (b BrandControllerImpl) GetBrandByNation(nationName string) []models.CarBrand {
	return b.Repo.GetBrandByNation(nationName)
}

func (b BrandControllerImpl) GetBrandByName(name string) []models.CarBrand {
	return b.Repo.GetBrandByName(name)
}

func (b BrandControllerImpl) AddBrand(name string, nation models.Nation) error {
	return b.Repo.AddNewBrand(models.CarBrand{
		Name:   name,
		Nation: nation,
	})
}
