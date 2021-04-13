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

	AddCar(modelName string, downloadUrl string, brand entities.CarBrand, categories []entities.CarCategory, year uint, drivetrain entities.Drivetrain, gearType entities.GearType) error
}

type TrackController interface {
	GetAllTracks() ([]entities.Track,error)
	GetTracksByNation(string) ([]entities.Track,error)
	GetTracksByLayoutType(string) ([]entities.Track,error)
	GetTracksByName(string) ([]entities.Track,error)

	AddTrack(name string, downloadUrl string, layouts []entities.Layout, location string, nation entities.Nation, year uint, tags []entities.TrackTag) error
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
