package repositories

import "github.com/davide/ModRepository/models"

type CarRepository interface {
	AddNewCar(car models.Car) error
	GetAllCars() []models.Car
}

type TrackRepository interface {
	AddNewTrack(track models.Track) error
	GetAllTracks() []models.Track
}

type NationRepository interface {
	AddNewNation(nation models.Nation) error
	GetAllNations() []models.Nation
}

type BrandRepository interface {
	AddNewBrand(brand models.CarBrand) error
	GetAllBrands() []models.CarBrand
}
