package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetAllCars() []models.Car {
	return c.Repo.GetAllCars()
}

func (c CarControllerImpl) GetCarsByNation(nationName string) []models.Car {
	return c.Repo.GetCarsByNation(nationName)
}

func (c CarControllerImpl) GetCarByModel(model string) []models.Car {
	return c.Repo.GetCarByModel(model)
}

func (c CarControllerImpl) GetCarsByBrand(brandName string) []models.Car {
	return c.Repo.GetCarsByBrand(brandName)
}

func (c CarControllerImpl) GetCarsByType(categoryName string) []models.Car {
	return c.Repo.GetCarsByType(categoryName)
}

func (c CarControllerImpl) AddCar(modelName string, downloadUrl string, brand models.CarBrand, categories []string) error {
	return c.Repo.AddNewCar(models.Car{ModelName: modelName, Brand: brand, Categories: categories, Mod: models.Mod{DownloadLink: downloadUrl}})
}
