package controllers

import "github.com/davide/ModRepository/models"

type CarController interface {
	GetAllCars() []models.Car
	GetCarsByNation(string) []models.Car
	GetCarByModel(string) []models.Car
	GetCarsByBrand(string) []models.Car
	GetCarsByType(string) []models.Car

	AddCar(modelName string, downloadUrl string, brand models.CarBrand, categories []string) error
}

type TrackController interface {
	GetAllTracks() []models.Track
	GetTracksByNation(string) []models.Track
	GetTracksByLayoutType(string) []models.Track
	GetTracksByName(string) []models.Track

	AddTrack(name string, downloadUrl string, layouts []models.Layout, location models.Location) error
}

type BrandController interface {
	GetAllBrands() []models.CarBrand
	GetBrandByNation(string) []models.CarBrand
	GetBrandByName(string) []models.CarBrand

	AddBrand(name string, nation models.Nation) error
}

type NationController interface {
	GetAllTracksNations() []models.Nation
	GetAllBrandsNations() []models.Nation
}
