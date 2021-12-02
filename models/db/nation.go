package db

type Nation struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Code   string         `gorm:"type:varchar(6)"`
	Brands []Manufacturer `gorm:"foreignKey:IdNation"`
	Tracks []Track        `gorm:"foreignKey:IdNation"`
}
