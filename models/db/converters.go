package db

import "github.com/davide/ModRepository/models/entities"

func CarFromEntity(car entities.Car) Car {

	return Car{
		DownloadLink: car.DownloadLink,
		ModelName:    car.ModelName,
		Brand:        car.Brand.Name,
		Categories:   allCarCategoryFromEntity(car.Categories),
		Year:         car.Year,
		Drivetrain:   string(car.Drivetrain),
		Transmission: string(car.Transmission),
		Premium:      car.Premium,
		Image:        car.Image,
		BHP:          car.BHP,
		Torque:       car.Torque,
		Weight:       car.Weight,
		TopSpeed:     car.TopSpeed,
		Author: car.Author.Name,
	}
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

func BrandFromEntity(brand entities.CarBrand) CarBrand {
	return CarBrand{
		Name:   brand.Name,
		Nation: brand.Nation.Name,
	}
}

func NationFromEntity(nation entities.Nation) Nation {
	return Nation{
		Name: nation.Name,
	}
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
		Image: track.Image,
		Author: track.Author.Name,
	}
}

func layoutFromEntity(layout entities.Layout, track string) Layout {
	return Layout{
		Name:     layout.Name,
		LengthM: layout.LengthM,
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