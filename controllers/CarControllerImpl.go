package controllers

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	repo repositories.CarRepository
}

func (c CarControllerImpl) getAllCars() []models.Car {
	return c.repo.GetAllCars()
}

func (c CarControllerImpl) getCarsByNation(s string) []models.Car {
	cars  := c.repo.GetAllCars()
	for i, car := range cars {
		if car.Brand.Nation.Name != s {
			cars = append(cars[:i], cars[i+1:]...)
		}
	}
	return cars
}

func (c CarControllerImpl) getCarByModel(s string) (error, models.Car) {
	for _, car := range c.repo.GetAllCars() {
		if car.ModelName == s {
			return nil, car
		}
	}
	return errors.New("car not found"), models.Car{}
}

func (c CarControllerImpl) getCarsByBrand(s string) []models.Car {
	cars  := c.repo.GetAllCars()
	for i, car := range cars {
		if car.Brand.Name != s {
			cars = append(cars[:i], cars[i+1:]...)
		}
	}
	return cars
}

func (c CarControllerImpl) getCarsByType(s string) []models.Car {
	cars  := c.repo.GetAllCars()
	for i, car := range cars {
		found := false
		for _, category := range car.Categories {
			if category == s {
				found = true
				break
			}
			if !found {
				cars = append(cars[:i], cars[i+1:]...)
			}
		}
	}
	return cars
}

func (c CarControllerImpl) addCar(modelName string, downloadUrl string, brand models.CarBrand, categories []string) error {
	return c.repo.AddNewCar(models.Car{ModelName: modelName, Brand: brand, Categories: categories, Mod: models.Mod{DownloadLink: downloadUrl}})
}
