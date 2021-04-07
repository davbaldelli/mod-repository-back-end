package controllers

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarController interface {
	GetAllCars() []entities.Car
	GetCarsByNation(string) []entities.Car
	GetCarByModel(string) []entities.Car
	GetCarsByBrand(string) []entities.Car
	GetCarsByType(string) []entities.Car

	AddCar(modelName string, downloadUrl string, brand entities.CarBrand, categories []entities.CarCategory) error
}

type TrackController interface {
	GetAllTracks() []entities.Track
	GetTracksByNation(string) []entities.Track
	GetTracksByLayoutType(string) []entities.Track
	GetTracksByName(string) []entities.Track

	AddTrack(name string, downloadUrl string, layouts []entities.Layout, location string, nation entities.Nation) error
}

type BrandController interface {
	GetAllBrands() []entities.CarBrand
	GetBrandByNation(string) []entities.CarBrand
	GetBrandByName(string) []entities.CarBrand

	AddBrand(name string, nation entities.Nation) error
}

type NationController interface {
	GetAllTracksNations() []entities.Nation
	GetAllBrandsNations() []entities.Nation
}
