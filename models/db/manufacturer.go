package db

import "github.com/davide/ModRepository/models/entities"

type Manufacturer struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Cars     []Car `gorm:"foreignKey:IdBrand"`
	IdNation uint
}

func (m Manufacturer) ToEntity(nation Nation) entities.CarBrand  {
	return entities.CarBrand{
		Name:   m.Name,
		Nation: entities.Nation{Name: nation.Name, Code: nation.Code},
	}
}
