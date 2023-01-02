package controllers

import (
	"github.com/davide/ModRepository/controllers/helpers"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	Repo repositories.CarRepository
}

func (c CarControllerImpl) GetAllCarCategories() ([]models.CarCategory, error) {
	return c.Repo.SelectAllCarCategories()
}

func (c CarControllerImpl) GetAllCars(role models.Role) ([]models.Car, error) {
	return c.Repo.SelectAllCars(helpers.IsPremium(role), helpers.IsAdmin(role))
}

func (c CarControllerImpl) AddCar(car *models.Car) error {
	return c.Repo.InsertCar(car)
}

func (c CarControllerImpl) UpdateCar(car models.Car) (bool, error) {
	return c.Repo.UpdateCar(car)
}
