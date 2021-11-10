package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetAllCarCategories(premium bool) ([]entities.CarCategory, error) {
	return c.Repo.SelectAllCarCategories(premium)
}

func (c CarControllerImpl) GetAllCars(premium bool) ([]entities.Car, error){
	return c.Repo.SelectAllCars(premium)
}

func (c CarControllerImpl) AddCar(car entities.Car) error {
	return c.Repo.InsertCar(car)
}
