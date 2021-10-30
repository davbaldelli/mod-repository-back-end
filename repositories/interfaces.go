package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	InsertCar(car entities.Car) error
	SelectAllCars(premium bool) ([]entities.Car, error)
	SelectCarsByNation(string, bool) ([]entities.Car, error)
	SelectCarsByModelName(string, bool) ([]entities.Car, error)
	SelectCarsByBrand(string, bool) ([]entities.Car, error)
	SelectCarsByType(string, bool) ([]entities.Car, error)
	SelectAllCarCategories(bool) ([]entities.CarCategory, error)
	SelectCarByModel(string) (entities.Car, error)
}

type TrackRepository interface {
	SelectAllTracks(bool) ([]entities.Track,error)
	SelectTracksByNation(string, bool) ([]entities.Track,error)
	SelectTracksByLayoutType(string, bool) ([]entities.Track,error)
	SelectTracksByName(string, bool) ([]entities.Track,error)
	SelectTracksByTag(entities.TrackTag, bool) ([]entities.Track,error)
	SelectTrackByName(string) (entities.Track, error)
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

type AuthorRepository interface {
	SelectAllAuthors() ([]entities.Author, error)
	SelectAllCarAuthors() ([]entities.Author, error)
	SelectAllTrackAuthors() ([]entities.Author, error)
}