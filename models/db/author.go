package db

import "github.com/davide/ModRepository/models/entities"

type Author struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Link   string
	Cars   []Car   `gorm:"foreignKey:IdAuthor"`
	Tracks []Track `gorm:"foreignKey:IdAuthor"`
}

func AuthorFromEntity(author entities.Author) Author {
	return Author{
		Name: author.Name,
		Link: author.Link,
	}
}
