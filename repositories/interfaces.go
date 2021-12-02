package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	InsertCar(car entities.Car) error
	SelectAllCars(premium bool) ([]entities.Car, error)
	SelectAllCarCategories() ([]entities.CarCategory, error)
	UpdateCar(car entities.Car) error
}

type LogRepository interface {
	SelectAllTrackLogs(bool) ([]entities.TrackLog, error)
	SelectAllCarLogs(bool) ([]entities.CarLog, error)
}

type TrackRepository interface {
	SelectAllTracks(bool) ([]entities.Track, error)
	InsertTrack(track entities.Track) error
	UpdateTrack(track entities.Track) error
}

type NationRepository interface {
	SelectAllBrandsNations() ([]entities.Nation, error)
	SelectAllTrackNations() ([]entities.Nation, error)
}

type BrandRepository interface {
	SelectAllBrands() ([]entities.CarBrand, error)
}

type UserRepository interface {
	Login(user entities.User) (entities.User, error)
	SignIn(user entities.User) (entities.User, error)
}

type AuthorRepository interface {
	SelectAllAuthors() ([]entities.Author, error)
	SelectAllCarAuthors() ([]entities.Author, error)
	SelectAllTrackAuthors() ([]entities.Author, error)
}
