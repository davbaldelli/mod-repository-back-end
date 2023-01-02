package controllers

import (
	"github.com/davide/ModRepository/models"
)

type CarController interface {
	GetAllCars(role models.Role) ([]models.Car, error)
	GetAllCarCategories() ([]models.CarCategory, error)
	AddCar(car *models.Car) error
	UpdateCar(car models.Car) (bool, error)
}

type TrackController interface {
	GetAllTracks(role models.Role) ([]models.Track, error)
	AddTrack(track *models.Track) error
	UpdateTrack(track models.Track) (bool, error)
}

type LogController interface {
	GetTrackLogs() ([]models.TrackLog, error)
	GetCarLogs() ([]models.CarLog, error)
}

type BrandController interface {
	GetAllBrands() ([]models.CarBrand, error)
}

type NationController interface {
	GetAllTracksNations() ([]models.Nation, error)
	GetAllBrandsNations() ([]models.Nation, error)
}

type UserController interface {
	Login(username string, password string) (models.User, error)
	SignIn(username string, password string, role models.Role) (models.User, error)
	UpdatePassword(username string, password string) error
}

type AuthorController interface {
	GetAllAuthors() ([]models.Author, error)
	GetAllCarAuthors() ([]models.Author, error)
	GetAllTrackAuthors() ([]models.Author, error)
}

type ServersController interface {
	GetAllServers() ([]models.Server, error)
	AddServer(server models.Server) error
	UpdateServer(server models.Server) error
	DeleteServer(server models.Server) error
}

type SkinController interface {
	SelectCarSkins(carId uint) ([]models.Skin, error)
	GetAllSkins() ([]models.Skin, error)
	AddSkin(skin models.Skin) error
	UpdateSkin(skin models.Skin) error
}

type FirebaseController interface {
	RegisterToTopic(token string, topic string) error
	NotifyCarAdded(car models.Car) error
	NotifyCarUpdated(car models.Car) error
	NotifyTrackUpdated(track models.Track) error
	NotifyTrackAdded(track models.Track) error
}

type DiscordBotController interface {
	NotifyCarAdded(car models.Car) error
	NotifyCarUpdated(car models.Car) error
	NotifyTrackUpdated(track models.Track) error
	NotifyTrackAdded(track models.Track) error
}
