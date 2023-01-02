package models

import (
	"github.com/davide/ModRepository/models"
)

type Manufacturer struct {
	Id       uint `gorm:"primaryKey"`
	Name     string
	Logo     string
	Cars     []Car `gorm:"foreignKey:IdBrand"`
	IdNation uint
}

func (m Manufacturer) ToEntity(nation Nation) models.CarBrand {
	return models.CarBrand{
		Name:   m.Name,
		Logo:   m.Logo,
		Nation: models.Nation{Name: nation.Name, Code: nation.Code},
	}
}

func ManufacturerFromEntity(brand models.CarBrand, idNation uint) Manufacturer {
	return Manufacturer{
		Name:     brand.Name,
		Logo:     brand.Logo,
		IdNation: idNation,
	}
}
