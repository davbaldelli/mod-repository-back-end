package models

import "github.com/davide/ModRepository/models/entities"

type Nation struct {
	Id     uint `gorm:"primaryKey"`
	Name   string
	Code   string         `gorm:"type:varchar(6)"`
	Brands []Manufacturer `gorm:"foreignKey:IdNation"`
	Tracks []Track        `gorm:"foreignKey:IdNation"`
}

func NationFromEntity(nation entities.Nation) Nation {
	return Nation{
		Name: nation.Name,
		Code: nation.Code,
	}
}
