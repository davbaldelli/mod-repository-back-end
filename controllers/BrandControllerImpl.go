package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	Repo repositories.BrandRepository
}

func (b BrandControllerImpl) GetAllBrands() []entities.CarBrand {
	return b.Repo.GetAllBrands()
}

func (b BrandControllerImpl) GetBrandByNation(nationName string) []entities.CarBrand {
	return b.Repo.GetBrandByNation(nationName)
}

func (b BrandControllerImpl) GetBrandByName(name string) []entities.CarBrand {
	return b.Repo.GetBrandByName(name)
}

