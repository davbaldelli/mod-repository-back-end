package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetAllCars() []entities.Car {
	return c.Repo.GetAllCars()
}

func (c CarControllerImpl) GetCarsByNation(nationName string) []entities.Car {
	return c.Repo.GetCarsByNation(nationName)
}

func (c CarControllerImpl) GetCarByModel(model string) []entities.Car {
	return c.Repo.GetCarByModel(model)
}

func (c CarControllerImpl) GetCarsByBrand(brandName string) []entities.Car {
	return c.Repo.GetCarsByBrand(brandName)
}

func (c CarControllerImpl) GetCarsByType(categoryName string) []entities.Car {
	return c.Repo.GetCarsByType(categoryName)
}

func (c CarControllerImpl) AddCar(modelName string, downloadUrl string, brand entities.CarBrand, categories []entities.CarCategory) error {
	return c.Repo.AddNewCar(entities.Car{ModelName: modelName, Brand: brand, Categories: categories, Mod: entities.Mod{DownloadLink: downloadUrl}})
}
