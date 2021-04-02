package controllers

import "github.com/davide/ModRepository/models"

type CarController interface {
	getAllCars() []models.Car
	getCarsByNation(string) []models.Car
	getCarByModel(string) (error, models.Car)
	getCarsByBrand(string) []models.Car
	getCarsByType(string) []models.Car

	addCar(modelName string, downloadUrl string, brand models.CarBrand, categories []string) error
}

type TrackController interface {
	getAllTracks() []models.Track
	getTracksByNation(string) []models.Track
	getTracksByLayoutType(string) []models.Track
	getTrackByName(string) (models.Track, error)

	addNewTrack(name string, downloadUrl string, layouts []models.Layout, locationName string, nation models.Nation) error
}

type BrandController interface {
	getAllBrands() []models.CarBrand
	getBrandByNation(string) []models.CarBrand
	getBrandByName(string) (models.CarBrand, error)

	addNewBrand(name string, nation models.Nation) error
}

type NationController interface {
	getAllNations() []models.Nation
	getNationByName(string) (models.Nation, error)

	addNewNation(string) error
}
