package db

import "github.com/davide/ModRepository/models/entities"

type Manufacturer struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Logo string
	Cars     []Car `gorm:"foreignKey:IdBrand"`
	IdNation uint
}

func (m Manufacturer) ToEntity(nation Nation) entities.CarBrand {
	return entities.CarBrand{
		Name:   m.Name,
		Logo: m.Logo,
		Nation: entities.Nation{Name: nation.Name, Code: nation.Code},
	}
}

func ManufacturerFromEntity(brand entities.CarBrand, idNation uint) Manufacturer{
	return Manufacturer{
		Name:     brand.Name,
		Logo:     brand.Logo,
		IdNation: idNation,
	}
}
