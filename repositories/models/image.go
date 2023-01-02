package models

import (
	"github.com/davide/ModRepository/models"
)

type Image struct {
	Id       uint `gorm:"primaryKey"`
	Url      string
	Favorite bool
}

func (i Image) toEntity() models.Image {
	return models.Image{
		Id:       i.Id,
		Url:      i.Url,
		Favorite: i.Favorite,
	}
}

func imageFromEntity(img models.Image) Image {
	return Image{
		Id:       img.Id,
		Url:      img.Url,
		Favorite: img.Favorite,
	}
}
