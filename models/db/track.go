package db

import "github.com/davide/ModRepository/models/entities"

type Track struct {
	ModModel
	Version      string
	DownloadLink string
	Source       string
	Personal     bool
	Name         string
	Layouts      []Layout `gorm:"foreignKey:IdTrack"`
	Location     string
	IdNation     uint
	Tags         []TrackTag `gorm:"foreignKey:IdTrack"`
	Year         uint
	Premium      bool
	Image        string
	IdAuthor     uint
	Rating       uint
}

type TrackMod struct {
	ModModel
	Version      string
	DownloadLink string
	Source       string
	Personal     bool
	Name         string
	Layouts      []Layout   `gorm:"foreignKey:IdTrack"`
	Tags         []TrackTag `gorm:"foreignKey:IdTrack"`
	Location     string
	Nation       string
	NationCode   string
	Year         uint
	Premium      bool
	Image        string
	Author       string
	AuthorLink   string
	Rating       uint
}

func (t TrackMod) ToEntity(userRole entities.Role) entities.Track {
	download := t.DownloadLink
	if (t.Premium && userRole == entities.Base) || (t.Personal && userRole != entities.Admin) {
		download = t.Source
	}
	return entities.Track{
		Mod: entities.Mod{
			Id:           t.Id,
			DownloadLink: download,
			Source:       t.Source,
			Premium:      t.Premium,
			Personal:     t.Personal,
			Image:        t.Image,
			Author: entities.Author{
				Name: t.Author,
				Link: t.AuthorLink,
			},
			Rating:    t.Rating,
			CreatedAt: t.CreatedAt,
			UpdatedAt: t.UpdatedAt,
			Version:   t.Version,
		},
		Name: t.Name,
		Layouts: mapLayouts(t.Layouts, func(layout Layout) entities.Layout {
			return layout.toEntity()
		}),
		Location: t.Location,
		Nation:   entities.Nation{Name: t.Nation, Code: t.NationCode},
		Tags: mapTags(t.Tags, func(tag TrackTag) entities.TrackTag {
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

func (t TrackTag) toEntity() entities.TrackTag {
	return entities.TrackTag(t.Tag)
}

type Layout struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	LengthM  float32
	Category string
	IdTrack  uint
}

func (l Layout) toEntity() entities.Layout {
	return entities.Layout{
		Name:     l.Name,
		LengthM:  l.LengthM,
		Category: entities.LayoutType(l.Category),
	}
}

func mapLayouts(vs []Layout, f func(layout Layout) entities.Layout) []entities.Layout {
	vsm := make([]entities.Layout, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func mapTags(vs []TrackTag, f func(tag TrackTag) entities.TrackTag) []entities.TrackTag {
	vsm := make([]entities.TrackTag, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func TrackFromEntity(track entities.Track, idNation uint, idAuthor uint) Track {
	var tags []TrackTag
	for _, tag := range track.Tags {
		tags = append(tags, TrackTag{Tag: string(tag)})
	}
	return Track{
		ModModel:     ModModel{Id: track.Id},
		DownloadLink: track.DownloadLink,
		Personal:     track.Personal,
		Name:         track.Name,
		Layouts:      allLayoutFromEntity(track.Layouts, idAuthor),
		Location:     track.Location,
		IdNation:     idNation,
		Tags:         tags,
		Year:         track.Year,
		Premium:      track.Premium,
		Image:        track.Image,
		IdAuthor:     idAuthor,
		Rating:       track.Rating,
		Version:      track.Version,
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
