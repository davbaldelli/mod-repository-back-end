package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetAllCarCategories() ([]entities.CarCategory, error) {
	return c.Repo.SelectAllCarCategories()
}

func (c CarControllerImpl) GetAllCars() ([]entities.Car, error){
	return c.Repo.SelectAllCars()
}

func (c CarControllerImpl) GetCarsByNation(nationName string) ([]entities.Car, error) {
	return c.Repo.SelectCarsByNation(nationName)
}

func (c CarControllerImpl) GetCarByModel(model string) ([]entities.Car, error) {
	return c.Repo.SelectCarsByModelName(model)
}

func (c CarControllerImpl) GetCarsByBrand(brandName string) ([]entities.Car, error) {
	return c.Repo.SelectCarsByBrand(brandName)
}

func (c CarControllerImpl) GetCarsByType(categoryName string) ([]entities.Car, error) {
	return c.Repo.SelectCarsByType(categoryName)
}

func (c CarControllerImpl) AddCar(car entities.Car) error {
	return c.Repo.InsertCar(car)
}
