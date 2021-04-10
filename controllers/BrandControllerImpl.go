package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	Repo repositories.BrandRepository
}

func (b BrandControllerImpl) GetAllBrands() ([]entities.CarBrand,error) {
	return b.Repo.SelectAllBrands()
}

func (b BrandControllerImpl) GetBrandsByNation(nationName string) ([]entities.CarBrand, error) {
	return b.Repo.SelectBrandsByNation(nationName)
}

func (b BrandControllerImpl) GetBrandsByName(name string) ([]entities.CarBrand, error) {
	return b.Repo.SelectBrandsByName(name)
}

