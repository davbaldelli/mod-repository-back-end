package controllers

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type CarControllerImpl struct {
	repo repositories.CarRepository
}

type carParamCondition func(models.Car) bool

func (c CarControllerImpl) getAllCars() []models.Car {
	return c.repo.GetAllCars()
}

func (c CarControllerImpl) getCarsByNation(nationName string) []models.Car {
	return c.carResearchByParam(func(car models.Car) bool { return car.Brand.Nation.Name == nationName })
}

func (c CarControllerImpl) getCarByModel(model string) (error, models.Car) {
	for _, car := range c.repo.GetAllCars() {
		if car.ModelName == model {
			return nil, car
		}
	}
	return errors.New("car" + model + " not found"), models.Car{}
}

func (c CarControllerImpl) getCarsByBrand(brandName string) []models.Car {
	return c.carResearchByParam(func(car models.Car) bool { return car.Brand.Name == brandName })
}

func (c CarControllerImpl) getCarsByType(categoryName string) []models.Car {
	var cars []models.Car
	for _, car := range c.repo.GetAllCars() {
		for _, category := range car.Categories {
			if category == categoryName {
				cars = append(cars, car)
				break
			}
		}
	}
	return cars
}

func (c CarControllerImpl) addCar(modelName string, downloadUrl string, brand models.CarBrand, categories []string) error {
	return c.repo.AddNewCar(models.Car{ModelName: modelName, Brand: brand, Categories: categories, Mod: models.Mod{DownloadLink: downloadUrl}})
}

func (c CarControllerImpl) carResearchByParam(cpc carParamCondition) []models.Car {
	var cars []models.Car
	for _, car := range c.repo.GetAllCars() {
		if cpc(car) {
			cars = append(cars, car)
		}
	}
	return cars
}
