package controllers

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarController interface {
	GetAllCars() ([]entities.Car, error)
	GetCarsByNation(string) ([]entities.Car, error)
	GetCarByModel(string) ([]entities.Car, error)
	GetCarsByBrand(string) ([]entities.Car, error)
	GetCarsByType(string) ([]entities.Car, error)
	GetAllCarCategories() ([]entities.CarCategory, error)

	AddCar(car entities.Car) error
}

type TrackController interface {
	GetAllTracks() ([]entities.Track,error)
	GetTracksByNation(string) ([]entities.Track,error)
	GetTracksByLayoutType(string) ([]entities.Track,error)
	GetTracksByName(string) ([]entities.Track,error)
	GetTracksByTag(tag entities.TrackTag) ([]entities.Track, error)

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
