package db

import "github.com/davide/ModRepository/models/entities"

type CarCategory struct {
	Name string `gorm:"primaryKey"`
}

func carCategoryFromEntity(category entities.CarCategory) CarCategory {
	return CarCategory{
		Name: category.Name,
	}
}

func allCarCategoryFromEntity(categories []entities.CarCategory) []CarCategory {
	var dbCats []CarCategory
	for _, cat := range categories {
		dbCats = append(dbCats, carCategoryFromEntity(cat))
	}
	return dbCats
}

type Car struct {
	DownloadLink string
	ModelName    string `gorm:"primaryKey"`
	Brand        string
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;"`
}

func CarFromEntity(car entities.Car) Car {

	return Car{
		DownloadLink: car.DownloadLink,
		ModelName:    car.ModelName,
		Brand:        car.Brand.Name,
		Categories:   allCarCategoryFromEntity(car.Categories),
	}
}

type CarBrand struct {
	Name   string `gorm:"primaryKey"`
	Cars   []Car  `gorm:"foreignKey:Brand"`
	Nation string
}

func BrandFromEntity(brand entities.CarBrand) CarBrand {
	return CarBrand{
		Name:   brand.Name,
		Nation: brand.Nation.Name,
	}
}

type Nation struct {
	Name   string     `gorm:"primaryKey"`
	Brands []CarBrand `gorm:"foreignKey:Nation"`
	Tracks []Track    `gorm:"foreignKey:Nation"`
}

func NationFromEntity(nation entities.Nation) Nation {
	return Nation{
		Name: nation.Name,
	}
}

type Track struct {
	DownloadLink string
	Name         string   `gorm:"primaryKey"`
	Layouts      []Layout `gorm:"foreignKey:Track"`
	Location     string
	Nation       string
}

func TrackFromEntity(track entities.Track) Track {
	return Track{
		DownloadLink: track.DownloadLink,
		Name:         track.Name,
		Layouts:      allLayoutFromEntity(track.Layouts, track.Name),
		Location:     track.Location,
		Nation:       track.Nation.Name,
	}
}

type Layout struct {
	Name     string `gorm:"primaryKey"`
	LengthKm float32
	Category string
	Track    string
}

func layoutFromEntity(layout entities.Layout, track string) Layout {
	return Layout{
		Name:     layout.Name,
		LengthKm: layout.LengthKm,
		Category: layout.Category.Name,
		Track:    track,
	}
}

func allLayoutFromEntity(layouts []entities.Layout, track string) []Layout {
	var dbLayouts []Layout
	for _, layout := range layouts {
		dbLayouts = append(dbLayouts, layoutFromEntity(layout, track))
	}
	return dbLayouts
}

type TrackCategory struct {
	Name    string   `gorm:"primaryKey"`
	Layouts []Layout `gorm:"foreignKey:Category"`
}
