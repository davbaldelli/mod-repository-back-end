package db

import "github.com/davide/ModRepository/models/entities"

type CarCategory struct {
	Name string `gorm:"primaryKey:not null"`
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
	Year 		uint
	Brand        string
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;"`
	GearType string
	Drivetrain string
	Premium bool
}

func CarFromEntity(car entities.Car) Car {

	return Car{
		DownloadLink: car.DownloadLink,
		ModelName:    car.ModelName,
		Brand:        car.Brand.Name,
		Categories:   allCarCategoryFromEntity(car.Categories),
		Year:         car.Year,
		Drivetrain: string(car.Drivetrain),
		GearType: string(car.GearType),
		Premium: car.Premium,
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
	Tags 		 []string
	Year 		 uint
	Premium 	 bool
}

func TrackFromEntity(track entities.Track) Track {
	var tags []string
	for _, tag := range track.Tags {
		tags = append(tags, string(tag))
	}
	return Track{
		DownloadLink: track.DownloadLink,
		Name:         track.Name,
		Layouts:      allLayoutFromEntity(track.Layouts, track.Name),
		Location:     track.Location,
		Nation:       track.Nation.Name,
		Tags: tags,
		Year: track.Year,
		Premium: track.Premium,
	}
}

type Layout struct {
	Name     string `gorm:"primaryKey"`
	LengthKm float32
	Category string
	Track    string `gorm:"primaryKey"`
}


func layoutFromEntity(layout entities.Layout, track string) Layout {
	return Layout{
		Name:     layout.Name,
		LengthKm: layout.LengthM,
		Category: string(layout.Category),
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

type User struct {
	Username string `gorm:"primaryKey"`
	Password string
	IsAdmin bool
}
