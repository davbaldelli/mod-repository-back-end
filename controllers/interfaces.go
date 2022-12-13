package controllers

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarController interface {
	GetAllCars(role entities.Role) ([]entities.Car, error)
	GetAllCarCategories() ([]entities.CarCategory, error)
	AddCar(car *entities.Car) error
	UpdateCar(car entities.Car) (bool, error)
}

type TrackController interface {
	GetAllTracks(role entities.Role) ([]entities.Track, error)
	AddTrack(track *entities.Track) error
	UpdateTrack(track entities.Track) (bool, error)
}

type LogController interface {
	GetTrackLogs() ([]entities.TrackLog, error)
	GetCarLogs() ([]entities.CarLog, error)
}

type BrandController interface {
	GetAllBrands() ([]entities.CarBrand, error)
}

type NationController interface {
	GetAllTracksNations() ([]entities.Nation, error)
	GetAllBrandsNations() ([]entities.Nation, error)
}

type UserController interface {
	Login(username string, password string) (entities.User, error)
	SignIn(username string, password string, role entities.Role) (entities.User, error)
	UpdatePassword(username string, password string) error
}

type AuthorController interface {
	GetAllAuthors() ([]entities.Author, error)
	GetAllCarAuthors() ([]entities.Author, error)
	GetAllTrackAuthors() ([]entities.Author, error)
}

type ServersController interface {
	GetAllServers() ([]entities.Server, error)
	AddServer(server entities.Server) error
	UpdateServer(server entities.Server) error
	DeleteServer(server entities.Server) error
}

type FirebaseController interface {
	RegisterToTopic(token string, topic string) error
	NotifyCarAdded(car entities.Car) error
	NotifyCarUpdated(car entities.Car) error
	NotifyTrackUpdated(track entities.Track) error
	NotifyTrackAdded(track entities.Track) error
}

type DiscordBotController interface {
	NotifyCarAdded(car entities.Car) error
	NotifyCarUpdated(car entities.Car) error
	NotifyTrackUpdated(track entities.Track) error
	NotifyTrackAdded(track entities.Track) error
}
