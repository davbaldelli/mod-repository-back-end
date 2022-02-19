package repositories

import (
	"github.com/davide/ModRepository/models/entities"
)

type CarRepository interface {
	InsertCar(car *entities.Car) error
	SelectAllCars(role entities.Role) ([]entities.Car, error)
	SelectAllCarCategories() ([]entities.CarCategory, error)
	UpdateCar(car entities.Car) (bool, error)
}

type LogRepository interface {
	SelectAllTrackLogs() ([]entities.TrackLog, error)
	SelectAllCarLogs() ([]entities.CarLog, error)
}

type TrackRepository interface {
	SelectAllTracks(role entities.Role) ([]entities.Track, error)
	InsertTrack(track *entities.Track) error
	UpdateTrack(track entities.Track) (bool, error)
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
