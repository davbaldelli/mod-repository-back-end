package models

import (
	"github.com/davide/ModRepository/models"
)

type Author struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Link   string
	Cars   []Car   `gorm:"foreignKey:IdAuthor"`
	Tracks []Track `gorm:"foreignKey:IdAuthor"`
}

func AuthorFromEntity(author models.Author) Author {
	return Author{
		Name: author.Name,
		Link: author.Link,
	}
}
