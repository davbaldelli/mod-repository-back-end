package entities

import (
	"github.com/davide/ModRepository/models"
)

type CarMods struct {
	ModModel
	ModelName    string `gorm:"column:model"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"foreignKey:CarId"`
	Transmission string
	Drivetrain   string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Author       string
	AuthorLink   string
	Nation       string
	NationCode   string
	NationFlag   string
	BrandLogo    string
	Images       []CarImage `gorm:"foreignKey:CarId"`
}

type Car struct {
	ModModel
	ModelName    string `gorm:"column:model"`
	Year         int
	IdBrand      uint
	Categories   []CarCategory `gorm:"foreignKey:CarId"`
	Transmission string
	Drivetrain   string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Images       []CarImage `gorm:"foreignKey:CarId"`
}

type CarCategory struct {
	Id       uint `gorm:"primaryKey"`
	Category string
	CarId    uint
}

type CarImage struct {
	Image
	CarId uint
}

func (CarCategory) TableName() string {
	return "car_categories"
}

func (CarImage) TableName() string {
	return "car_images"
}

func (cat CarCategory) toEntity() models.CarCategory {
	return models.CarCategory{Name: models.CarType(cat.Category)}
}

func (c CarMods) ToEntity(premium bool, admin bool) models.Car {
	download := c.DownloadLink
	if (c.Premium && !premium) || (c.Personal && !admin) {
		download = c.Source
	}
	return models.Car{
		Mod: models.Mod{
			Id:           c.Id,
			DownloadLink: download,
			Source:       c.Source,
			Premium:      c.Premium,
			Personal:     c.Personal,
			Images:       allCarImagesToEntity(c.Images),
			Author: models.Author{
				Name: c.Author,
				Link: c.AuthorLink,
			},
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
			Rating:    c.Rating,
			Version:   c.Version,
			Official:  c.Official,
		},
		Brand: models.CarBrand{
			Name: c.Brand,
			Nation: models.Nation{
				Name: c.Nation,
				Code: c.NationCode,
				Flag: c.NationFlag,
			},
		},
		ModelName: c.ModelName,
		Categories: mapCategories(c.Categories, func(category CarCategory) models.CarCategory {
			return category.toEntity()
		}),
		Drivetrain:   models.Drivetrain(c.Drivetrain),
		Transmission: models.Transmission(c.Transmission),
		Year:         c.Year,
		Torque:       c.Torque,
		TopSpeed:     c.TopSpeed,
		Weight:       c.Weight,
		BHP:          c.BHP,
	}
}

func mapCategories(vs []CarCategory, f func(category CarCategory) models.CarCategory) []models.CarCategory {
	vsm := make([]models.CarCategory, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func CarFromEntity(car models.Car, idBrand uint, idAuthor uint) Car {
	return Car{
		ModModel: ModModel{
			Id:           car.Id,
			Rating:       car.Rating,
			Version:      car.Version,
			DownloadLink: car.DownloadLink,
			Source:       car.Source,
			Premium:      car.Premium,
			Personal:     car.Personal,
			IdAuthor:     idAuthor,
			Official:     car.Official,
		},

		ModelName:    car.ModelName,
		IdBrand:      idBrand,
		Categories:   allCarCategoryFromEntity(car.Categories, car.Id),
		Year:         int(car.Year),
		Drivetrain:   string(car.Drivetrain),
		Transmission: string(car.Transmission),
		BHP:          car.BHP,
		Torque:       car.Torque,
		Weight:       car.Weight,
		TopSpeed:     car.TopSpeed,
		Images:       allCarImagesFromEntity(car.Images, car.Id),
	}
}

func carCategoryFromEntity(category models.CarCategory, id uint) CarCategory {
	return CarCategory{
		CarId:    id,
		Category: string(category.Name),
	}
}

func allCarCategoryFromEntity(categories []models.CarCategory, id uint) []CarCategory {
	var dbCats []CarCategory
	for _, cat := range categories {
		dbCats = append(dbCats, carCategoryFromEntity(cat, id))
	}
	return dbCats
}

func (i CarImage) toEntity() models.Image {
	return i.Image.toEntity()
}
func allCarImagesToEntity(dbImages []CarImage) []models.Image {
	var images []models.Image
	for _, dbImage := range dbImages {
		images = append(images, dbImage.toEntity())
	}
	return images
}

func carImageFromEntity(image models.Image, id uint) CarImage {
	return CarImage{
		Image: imageFromEntity(image),
		CarId: id,
	}
}

func allCarImagesFromEntity(images []models.Image, id uint) []CarImage {
	var dbImages []CarImage
	for _, image := range images {
		dbImages = append(dbImages, carImageFromEntity(image, id))
	}
	return dbImages
}
