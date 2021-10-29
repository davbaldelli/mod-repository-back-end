package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetCarByModel(model string) (entities.Car, error) {
	return c.Repo.SelectCarByModel(model)
}

func (c CarControllerImpl) GetAllCarCategories() ([]entities.CarCategory, error) {
	return c.Repo.SelectAllCarCategories()
}

func (c CarControllerImpl) GetAllCars(premium bool) ([]entities.Car, error){
	return c.Repo.SelectAllCars(premium)
}

func (c CarControllerImpl) GetCarsByNation(nationName string, premium bool) ([]entities.Car, error) {
	return c.Repo.SelectCarsByNation(nationName, premium)
}

func (c CarControllerImpl) GetCarsByModel(model string, premium bool) ([]entities.Car, error) {
	return c.Repo.SelectCarsByModelName(model, premium)
}

func (c CarControllerImpl) GetCarsByBrand(brandName string, premium bool) ([]entities.Car, error) {
	return c.Repo.SelectCarsByBrand(brandName, premium)
}

func (c CarControllerImpl) GetCarsByType(categoryName string, premium bool) ([]entities.Car, error) {
	return c.Repo.SelectCarsByType(categoryName, premium)
}

func (c CarControllerImpl) AddCar(car entities.Car) error {
	return c.Repo.InsertCar(car)
}
