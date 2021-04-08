package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	AddNewCar(car entities.Car) error
	GetAllCars() []entities.Car
	GetCarsByNation(string) []entities.Car
	GetCarByModel(string) []entities.Car
	GetCarsByBrand(string) []entities.Car
	GetCarsByType(string) []entities.Car
}

type TrackRepository interface {
	GetAllTracks() []entities.Track
	GetTracksByNation(string) []entities.Track
	GetTracksByLayoutType(string) []entities.Track
	GetTracksByName(string) []entities.Track
	AddNewTrack(track entities.Track) error
}

type NationRepository interface {
	GetAllBrandsNations() []entities.Nation
	GetAllTrackNations() []entities.Nation
}

type BrandRepository interface {
	GetAllBrands() []entities.CarBrand
	GetBrandByNation(string) []entities.CarBrand
	GetBrandByName(string) []entities.CarBrand
}
