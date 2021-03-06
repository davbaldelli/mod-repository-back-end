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

func (c CarControllerImpl) GetAllCars(role entities.Role) ([]entities.Car, error) {
	return c.Repo.SelectAllCars(role)
}

func (c CarControllerImpl) AddCar(car *entities.Car) error {
	return c.Repo.InsertCar(car)
}

func (c CarControllerImpl) UpdateCar(car entities.Car) (bool, error) {
	return c.Repo.UpdateCar(car)
}
