package db

import "github.com/davide/ModRepository/models/entities"

type Image struct {
	Id  uint `gorm:"primaryKey"`
	Url string
}

func (i Image) toEntity() entities.Image {
	return entities.Image{
		Id:  i.Id,
		Url: i.Url,
	}
}

func imageFromEntity(img entities.Image) Image {
	return Image{
		Id:  img.Id,
		Url: img.Url,
	}
}

func allImagesToEntity(dbImages []Image) []entities.Image {
	var images []entities.Image
	for _, dbImage := range dbImages {
		images = append(images, dbImage.toEntity())
	}
	return images
}
