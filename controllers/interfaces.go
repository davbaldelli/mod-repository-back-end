package controllers

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarController interface {
	GetAllCars(premium bool) ([]entities.Car, error)
	GetCarsByNation(string, bool) ([]entities.Car, error)
	GetCarsByModel(string, bool) ([]entities.Car, error)
	GetCarsByBrand(string, bool) ([]entities.Car, error)
	GetCarsByType(string, bool) ([]entities.Car, error)
	GetAllCarCategories() ([]entities.CarCategory, error)
	GetCarByModel(string) (entities.Car, error)

	AddCar(car entities.Car) error
}

type TrackController interface {
	GetAllTracks() ([]entities.Track,error)
	GetTracksByNation(string) ([]entities.Track,error)
	GetTracksByLayoutType(string) ([]entities.Track,error)
	GetTracksByName(string) ([]entities.Track,error)
	GetTracksByTag(tag entities.TrackTag) ([]entities.Track, error)
	GetTrackByName(string) (entities.Track, error)

	AddTrack(track entities.Track) error
}

type BrandController interface {
	GetAllBrands() ([]entities.CarBrand, error)
	GetBrandsByNation(string) ([]entities.CarBrand, error)
	GetBrandsByName(string) ([]entities.CarBrand, error)
}

type NationController interface {
	GetAllTracksNations() ([]entities.Nation, error)
	GetAllBrandsNations() ([]entities.Nation, error)
}

type LoginController interface {
	Login(username string, password string) (entities.User,error)
}

type AuthorController interface {
	GetAllAuthors() ([]entities.Author, error)
	GetAllCarAuthors() ([]entities.Author, error)
	GetAllTrackAuthors() ([]entities.Author, error)
}
