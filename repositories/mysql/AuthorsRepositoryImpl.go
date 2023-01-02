package mysql

import (
	"github.com/davide/ModRepository/models"
	"gorm.io/gorm"
)

type AuthorsRepositoryImpl struct {
	Db *gorm.DB
}

func (a AuthorsRepositoryImpl) SelectAllAuthors() ([]models.Author, error) {
	var authors []models.Author
	if result := a.Db.Order("authors.name ASC").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}

func (a AuthorsRepositoryImpl) SelectAllCarAuthors() ([]models.Author, error) {
	var authors []models.Author
	if result := a.Db.Order("authors.name ASC").Distinct().Joins("join cars on id_author = authors.id").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}

func (a AuthorsRepositoryImpl) SelectAllTrackAuthors() ([]models.Author, error) {
	var authors []models.Author
	if result := a.Db.Order("authors.name ASC").Distinct().Joins("join tracks on id_author = authors.id").Find(&authors); result.Error != nil {
		return authors, result.Error
	} else {
		return authors, nil
	}
}
