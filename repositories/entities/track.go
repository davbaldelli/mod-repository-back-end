package entities

import (
	"github.com/davide/ModRepository/models"
)

type Track struct {
	ModModel
	Name     string
	Layouts  []Layout `gorm:"foreignKey:IdTrack"`
	Location string
	IdNation uint
	Tags     []TrackTag `gorm:"foreignKey:IdTrack"`
	Year     uint
	Images   []TrackImage `gorm:"foreignKey:TrackId"`
}

type TrackMod struct {
	ModModel
	Name       string
	Layouts    []Layout   `gorm:"foreignKey:IdTrack"`
	Tags       []TrackTag `gorm:"foreignKey:IdTrack"`
	Location   string
	Nation     string
	NationCode string
	NationFlag string
	Year       uint
	Author     string
	AuthorLink string
	Images     []TrackImage `gorm:"foreignKey:TrackId"`
}

type TrackImage struct {
	Image
	TrackId uint
}

func (t TrackMod) ToEntity(premium bool, admin bool) models.Track {
	download := t.DownloadLink
	if (t.Premium && !premium) || (t.Personal && !admin) {
		download = t.Source
	}
	return models.Track{
		Mod: models.Mod{
			Id:           t.Id,
			DownloadLink: download,
			Source:       t.Source,
			Premium:      t.Premium,
			Personal:     t.Personal,
			Images:       allTrackImagesToEntity(t.Images),
			Author: models.Author{
				Name: t.Author,
				Link: t.AuthorLink,
			},
			Rating:    t.Rating,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			Version:   t.Version,
			Official:  t.Official,
		},
		Name: t.Name,
		Layouts: mapLayouts(t.Layouts, func(layout Layout) models.Layout {
			return layout.toEntity()
		}),
		Location: t.Location,
		Nation:   models.Nation{Name: t.Nation, Code: t.NationCode, Flag: t.NationFlag},
		Tags: mapTags(t.Tags, func(tag TrackTag) models.TrackTag {
			return tag.toEntity()
		}),
		Year: t.Year,
	}
}

type TrackTag struct {
	Id      uint `gorm:"primarykey"`
	IdTrack uint
	Tag     string
}

func (t TrackTag) toEntity() models.TrackTag {
	return models.TrackTag(t.Tag)
}

type Layout struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	LengthM  float32
	Category string
	IdTrack  uint
}

func (l Layout) toEntity() models.Layout {
	return models.Layout{
		Name:     l.Name,
		LengthM:  l.LengthM,
		Category: models.LayoutType(l.Category),
	}
}

func mapLayouts(vs []Layout, f func(layout Layout) models.Layout) []models.Layout {
	vsm := make([]models.Layout, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func mapTags(vs []TrackTag, f func(tag TrackTag) models.TrackTag) []models.TrackTag {
	vsm := make([]models.TrackTag, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TrackFromEntity(track models.Track, idNation uint, idAuthor uint) Track {
	var tags []TrackTag
	for _, tag := range track.Tags {
		tags = append(tags, TrackTag{Tag: string(tag)})
	}
	return Track{
		ModModel: ModModel{
			Id:           track.Id,
			Rating:       track.Rating,
			Version:      track.Version,
			DownloadLink: track.DownloadLink,
			Source:       track.Source,
			Premium:      track.Premium,
			Personal:     track.Personal,
			IdAuthor:     idAuthor,
			Official:     track.Official,
		},
		Name:     track.Name,
		Layouts:  allLayoutFromEntity(track.Layouts, idAuthor),
		Location: track.Location,
		IdNation: idNation,
		Tags:     tags,
		Year:     track.Year,
		Images:   allTrackImagesFromEntity(track.Images, track.Id),
	}
}

func layoutFromEntity(layout models.Layout, idTrack uint) Layout {
	return Layout{
		Name:     layout.Name,
		LengthM:  layout.LengthM,
		Category: string(layout.Category),
		IdTrack:  idTrack,
	}
}

func allLayoutFromEntity(layouts []models.Layout, track uint) []Layout {
	var dbLayouts []Layout
	for _, layout := range layouts {
		dbLayouts = append(dbLayouts, layoutFromEntity(layout, track))
	}
	return dbLayouts
}

func (i TrackImage) toEntity() models.Image {
	return i.Image.toEntity()
}
func allTrackImagesToEntity(dbImages []TrackImage) []models.Image {
	var images []models.Image
	for _, dbImage := range dbImages {
		images = append(images, dbImage.toEntity())
	}
	return images
}

func trackImageFromEntity(image models.Image, id uint) TrackImage {
	return TrackImage{
		Image:   imageFromEntity(image),
		TrackId: id,
	}
}

func allTrackImagesFromEntity(images []models.Image, id uint) []TrackImage {
	var dbImages []TrackImage
	for _, image := range images {
		dbImages = append(dbImages, trackImageFromEntity(image, id))
	}
	return dbImages
}
