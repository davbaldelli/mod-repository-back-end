package postgresrepo

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrackRepositoryImpl struct {
	Db *gorm.DB
}

type selectFromTrackQuery func(*[]db.TrackMod) *gorm.DB

func trackToEntity(dbTrack db.TrackMod) entities.Track {
	var tags []entities.TrackTag
	for _, tag := range dbTrack.Tags {
		tags = append(tags, entities.TrackTag(tag))
	}
	return entities.Track{
		Mod:      entities.Mod{
			DownloadLink: dbTrack.DownloadLink,
			Premium: dbTrack.Premium,
			Image: dbTrack.Image,
			Author: entities.Author{
				Name: dbTrack.Author,
				Link: dbTrack.AuthorLink,
			},
		},
		Name:     dbTrack.Name,
		Layouts:  allLayoutsToEntity(dbTrack.Layouts),
		Location: dbTrack.Location,
		Nation:   entities.Nation{Name: dbTrack.Nation},
		Tags: tags,
		Year: dbTrack.Year,
	}
}

func allLayoutsToEntity(dbLayouts []db.Layout) []entities.Layout{
	var layouts []entities.Layout
	for _,dbLayout := range dbLayouts{
		layouts = append(layouts, entities.Layout{
			Name:     dbLayout.Name,
			LengthM:  dbLayout.LengthM,
			Category: entities.LayoutType(dbLayout.Category),
		})
	}
	return layouts
}

func (t TrackRepositoryImpl) SelectTrackByName(name string) (entities.Track, error) {
	track := db.TrackMod{Name: name}
	if result := t.Db.Preload("Layouts").First(&track); result.Error != nil{
		return entities.Track{}, result.Error
	}

	author := db.Author{Name: track.Author}
	if result := t.Db.First(&author); result.Error != nil{
		return entities.Track{}, result.Error
	}

	return trackToEntity(track), nil
}

func (t TrackRepositoryImpl) SelectAllTracks() ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.TrackMod) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(tracks)
	})
}

func (t TrackRepositoryImpl) SelectTracksByNation(nation string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.TrackMod) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"nation = ?",nation)
	})
}

func (t TrackRepositoryImpl) SelectTracksByLayoutType(category string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.TrackMod) *gorm.DB {
		return t.Db.Order("name ASC").Distinct().Preload("Layouts").Joins("join layouts on layouts.track = track_mods.name").Where("layouts.category = ?", category).Find(&tracks)
	})
}

func (t TrackRepositoryImpl) SelectTracksByName(name string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.TrackMod) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"LOWER(track_mods.name) LIKE LOWER(?)","%"+name+"%")
	})
}

func (t TrackRepositoryImpl) SelectTrackByTag(tag entities.TrackTag) ([]entities.Track,error){
	return selectTracksWithQuery(func(tracks *[]db.TrackMod) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks," ? = ANY (tags)", tag)
	})
}

func (t TrackRepositoryImpl) InsertTrack(track entities.Track) error {
	dbTrack := db.TrackFromEntity(track)
	dbNation := db.NationFromEntity(track.Nation)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&track.Author); res.Error != nil {
		return res.Error
	}

	if res := t.Db.Model(&db.Nation{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := t.Db.Create(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}

func selectTracksWithQuery(query selectFromTrackQuery) ([]entities.Track, error) {
	var dbTracks []db.TrackMod
	var tracks []entities.Track
	if result := query(&dbTracks); result.Error != nil {
		return nil,result.Error
	}
	for _, dbTrack := range dbTracks {
		var tags []entities.TrackTag
		for _, tag := range dbTrack.Tags {
			tags = append(tags, entities.TrackTag(tag))
		}
		tracks = append(tracks, trackToEntity(dbTrack))
	}
	return tracks,nil
}
