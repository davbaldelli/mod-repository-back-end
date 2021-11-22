package db

import (
	"github.com/davide/ModRepository/models/entities"
)

func CarFromEntity(car entities.Car, idBrand uint, idAuthor uint) Car {
	return Car{
		ModModel : ModModel{Id: car.Id},
		DownloadLink: car.DownloadLink,
		ModelName:    car.ModelName,
		IdBrand:      idBrand,
		Categories:   allCarCategoryFromEntity(car.Categories),
		Year: int(car.Year),
		Drivetrain:   string(car.Drivetrain),
		Transmission: string(car.Transmission),
		Premium:      car.Premium,
		Image:        car.Image,
		BHP:          car.BHP,
		Torque:       car.Torque,
		Weight:       car.Weight,
		TopSpeed:     car.TopSpeed,
		IdAuthor:     idAuthor,
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

func BrandFromEntity(brand entities.CarBrand, idNation uint) Manufacturer {
	return Manufacturer{
		Name:     brand.Name,
		IdNation: idNation,
	}
}

func NationFromEntity(nation entities.Nation) Nation {
	return Nation{
		Name: nation.Name,
		Code: nation.Code,
	}
}

func TrackFromEntity(track entities.Track, idNation uint, idAuthor uint) Track {
	var tags []TrackTag
	for _, tag := range track.Tags {
		tags = append(tags, TrackTag{Tag: string(tag)})
	}
	return Track{
		ModModel: ModModel{Id: track.Id},
		DownloadLink: track.DownloadLink,
		Name:         track.Name,
		Layouts:      allLayoutFromEntity(track.Layouts, idAuthor),
		Location:     track.Location,
		IdNation:     idNation,
		Tags: tags,
		Year:     track.Year,
		Premium:  track.Premium,
		Image:    track.Image,
		IdAuthor: idAuthor,
	}
}

func layoutFromEntity(layout entities.Layout, idTrack uint) Layout {
	return Layout{
		Name:     layout.Name,
		LengthM:  layout.LengthM,
		Category: string(layout.Category),
		IdTrack:  idTrack,
	}
}

func allLayoutFromEntity(layouts []entities.Layout, track uint) []Layout {
	var dbLayouts []Layout
	for _, layout := range layouts {
		dbLayouts = append(dbLayouts, layoutFromEntity(layout, track))
	}
	return dbLayouts
}
