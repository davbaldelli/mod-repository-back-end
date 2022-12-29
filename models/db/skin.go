package db

import "github.com/davide/ModRepository/models/entities"

type Skin struct {
	Id           uint `gorm:"primaryKey"`
	Name         string
	DownloadLink string
	ImageUrl     string
	CarId        uint
}

type Skins []Skin

func (s Skin) toEntity() entities.Skin {
	return entities.Skin{
		Id:           s.Id,
		Name:         s.Name,
		DownloadLink: s.DownloadLink,
		ImageUrl:     s.ImageUrl,
		CarId:        s.CarId,
	}
}

func (s Skins) toEntities() []entities.Skin {
	var skins []entities.Skin
	for _, dbSkin := range s {
		skins = append(skins, dbSkin.toEntity())
	}
	return skins
}

func SkinFromEntity(skin entities.Skin) Skin {
	return Skin{
		Id:           skin.Id,
		Name:         skin.Name,
		DownloadLink: skin.DownloadLink,
		ImageUrl:     skin.ImageUrl,
		CarId:        skin.CarId,
	}
}

func SkinsFromEntities(skins []entities.Skin) []Skin {
	var dbSkins []Skin
	for _, skin := range skins {
		dbSkins = append(dbSkins, SkinFromEntity(skin))
	}
	return dbSkins
}
