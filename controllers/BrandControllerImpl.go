package controllers

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type BrandControllerImpl struct {
	repo repositories.BrandRepository
}

type paramCondition func(models.CarBrand) bool

func (b BrandControllerImpl) getAllBrands() []models.CarBrand {
	return b.repo.GetAllBrands()
}

func (b BrandControllerImpl) getBrandByNation(nationName string) []models.CarBrand {
	return b.brandResearchByParam(
		func(brand models.CarBrand) bool { return brand.Name == nationName })
}

func (b BrandControllerImpl) getBrandByName(name string) (models.CarBrand, error) {
	for _, brand := range b.repo.GetAllBrands() {
		if brand.Name == name {
			return brand, nil
		}
	}
	return models.CarBrand{}, errors.New("brand" + name + "not found")
}

func (b BrandControllerImpl) addNewBrand(name string, nation models.Nation) error {
	return b.repo.AddNewBrand(models.CarBrand{
		Name:   name,
		Nation: nation,
	})
}

func (b BrandControllerImpl) brandResearchByParam(pc paramCondition) []models.CarBrand {
	var brands []models.CarBrand
	for _, brand := range b.repo.GetAllBrands() {
		if pc(brand) {
			brands = append(brands, brand)
		}
	}
	return brands
}
