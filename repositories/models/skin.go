package models

import (
	"github.com/davide/ModRepository/models"
)

type Skin struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	DownloadLink string
	ImageUrl     string
	CarId        uint
}

type Skins []Skin

func (s Skin) toEntity() models.Skin {
	return models.Skin{
		Id:           s.Id,
		Name:         s.Name,
		DownloadLink: s.DownloadLink,
		ImageUrl:     s.ImageUrl,
		CarId:        s.CarId,
	}
}

func (s Skins) toEntities() []models.Skin {
	var skins []models.Skin
	for _, dbSkin := range s {
		skins = append(skins, dbSkin.toEntity())
	}
	return skins
}
