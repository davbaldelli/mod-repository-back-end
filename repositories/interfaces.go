package repositories

import (
	"github.com/davide/ModRepository/models"
)

type CarRepository interface {
	InsertCar(car *models.Car) error
	SelectAllCars(premium bool, admin bool) ([]models.Car, error)
	SelectAllCarCategories() ([]models.CarCategory, error)
	UpdateCar(car models.Car) (bool, error)
}

type LogRepository interface {
	SelectAllTrackLogs() ([]models.TrackLog, error)
	SelectAllCarLogs() ([]models.CarLog, error)
}

type TrackRepository interface {
	SelectAllTracks(premium bool, admin bool) ([]models.Track, error)
	InsertTrack(track *models.Track) error
	UpdateTrack(track models.Track) (bool, error)
}

type NationRepository interface {
	SelectAllBrandsNations() ([]models.Nation, error)
	SelectAllTrackNations() ([]models.Nation, error)
}

type BrandRepository interface {
	SelectAllBrands() ([]models.CarBrand, error)
}

type UserRepository interface {
	Login(user models.User) (models.User, error)
	SignIn(user models.User) (models.User, error)
	UpdatePassword(string, string) error
}

type AuthorRepository interface {
	SelectAllAuthors() ([]models.Author, error)
	SelectAllCarAuthors() ([]models.Author, error)
	SelectAllTrackAuthors() ([]models.Author, error)
}

type ServersRepository interface {
	GetAllServers() ([]models.Server, error)
	UpdateServer(server models.Server) error
	AddServer(server models.Server) error
	DeleteServer(server models.Server) error
}

type SkinRepository interface {
	SelectCarSkins(carId uint) ([]models.Skin, error)
	GetAllSkins() ([]models.Skin, error)
	AddSkin(skin models.Skin) error
	UpdateSkin(skin models.Skin) error
}
