package mysql

import (
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type AuthorsRepositoryImpl struct {
	Db *gorm.DB
}

func (a AuthorsRepositoryImpl) SelectAllAuthors() ([]entities.Author, error) {
	var authors []entities.Author
	if result := a.Db.Order("name ASC").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}

func (a AuthorsRepositoryImpl) SelectAllCarAuthors() ([]entities.Author, error) {
	var authors []entities.Author
	if result := a.Db.Order("name ASC").Distinct().Joins("join cars on id_author = authors.id").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}

func (a AuthorsRepositoryImpl) SelectAllTrackAuthors() ([]entities.Author, error) {
	var authors []entities.Author
	if result := a.Db.Order("name ASC").Distinct().Joins("join tracks on id_author = authors.id").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}
