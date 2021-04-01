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
	getTrackByName(string) []models.Track

	addNewTrack(name string, downloadUrl string, layouts []models.Layout, locationName string, nation models.Nation)
}

type BrandController interface {
	getAllBrands(string) []models.CarBrand
	getBrandByNation(string) []models.CarBrand
	getBrandByName(string) models.CarBrand

	addNewBrand(name string, nation models.Nation) error
}

type NationController interface {
	getAllNations() []models.Nation
	getNationByName(string) models.Nation

	addNewNation(string)
}