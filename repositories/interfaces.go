package repositories

import "github.com/davide/ModRepository/models"

type CarRepository interface {
	AddNewCar(car models.Car) error
	GetAllCars() []models.Car
}

type TrackRepository interface {
	addNewTrack(track models.Track) error
	getAllTracks() []models.Track
}

type NationRepository interface {
	addNewNation(nation models.Nation) error
	getAllNations() []models.Nation
}

type BrandRepository interface {
	addNewBrand(brand models.CarBrand)
	getAllBrands() []models.CarBrand
}