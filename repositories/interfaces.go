package repositories

import "github.com/davide/ModRepository/models"

type CarRepository interface {
	AddNewCar(car models.Car) error
	GetAllCars() []models.Car
	GetCarsByNation(string) []models.Car
	GetCarByModel(string) []models.Car
	GetCarsByBrand(string) []models.Car
	GetCarsByType(string) []models.Car
}

type TrackRepository interface {
	GetAllTracks() []models.Track
	GetTracksByNation(string) []models.Track
	GetTracksByLayoutType(string) []models.Track
	GetTracksByName(string) []models.Track
	AddNewTrack(track models.Track) error
}

type NationRepository interface {
	GetAllBrandsNations() []models.Nation
	GetAllTrackNations() []models.Nation
}

type BrandRepository interface {
	AddNewBrand(brand models.CarBrand) error
	GetAllBrands() []models.CarBrand
	GetBrandByNation(string) []models.CarBrand
	GetBrandByName(string) []models.CarBrand
}
