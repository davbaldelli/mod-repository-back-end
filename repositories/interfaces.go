package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	InsertCar(car entities.Car) error
	SelectAllCars() ([]entities.Car, error)
	SelectCarsByNation(string) ([]entities.Car, error)
	SelectCarsByModelName(string) ([]entities.Car, error)
	SelectCarsByBrand(string) ([]entities.Car, error)
	SelectCarsByType(string) ([]entities.Car, error)
	SelectAllCarCategories() ([]entities.CarCategory, error)
}

type TrackRepository interface {
	SelectAllTracks() ([]entities.Track,error)
	SelectTracksByNation(string) ([]entities.Track,error)
	SelectTracksByLayoutType(string) ([]entities.Track,error)
	SelectTracksByName(string) ([]entities.Track,error)
	SelectAllTrackCategories() ([]entities.TrackCategory, error)
	InsertTrack(track entities.Track) error
}

type NationRepository interface {
	SelectAllBrandsNations() ([]entities.Nation, error)
	SelectAllTrackNations() ([]entities.Nation, error)
}

type BrandRepository interface {
	SelectAllBrands() ([]entities.CarBrand, error)
	SelectBrandsByNation(string) ([]entities.CarBrand, error)
	SelectBrandsByName(string) ([]entities.CarBrand, error)
}

type UserRepository interface {
	Login(user entities.User) (entities.User, error)
}