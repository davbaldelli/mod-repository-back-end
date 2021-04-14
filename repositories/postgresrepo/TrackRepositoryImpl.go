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

type selectFromTrackQuery func(*[]db.Track) *gorm.DB




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

func (t TrackRepositoryImpl) SelectAllTracks() ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(tracks)
	})
}

func (t TrackRepositoryImpl) SelectTracksByNation(nation string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"nation = ?",nation)
	})
}

func (t TrackRepositoryImpl) SelectTracksByLayoutType(category string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Joins("join layouts on layouts.track = tracks.name").Where("layouts.category = ?", category).Find(&tracks)
	})
}

func (t TrackRepositoryImpl) SelectTracksByName(name string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"LOWER(tracks.name) LIKE LOWER(?)","%"+name+"%")
	})
}

func (t TrackRepositoryImpl) InsertTrack(track entities.Track) error {
	dbTrack := db.TrackFromEntity(track)
	dbNation := db.NationFromEntity(track.Nation)

	if res := t.Db.Model(&db.Nation{}).Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := t.Db.Model(&db.Track{}).Create(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}

func selectTracksWithQuery(query selectFromTrackQuery) ([]entities.Track, error) {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := query(&dbTracks); result.Error != nil {
		return nil,result.Error
	}
	for _, dbTrack := range dbTracks {
		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{DownloadLink: dbTrack.DownloadLink, Premium: dbTrack.Premium},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
		})
	}
	return tracks,nil
}
