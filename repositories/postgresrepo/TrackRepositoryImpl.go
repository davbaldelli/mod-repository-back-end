package postgresrepo

import (
	"fmt"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrackRepositoryImpl struct {
	Db *gorm.DB
}

func (t TrackRepositoryImpl) GetAllTracks() []entities.Track {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := t.Db.Preload("Layouts").Find(&dbTracks); result.Error != nil {
		//return error
	}
	for _, dbTrack := range dbTracks {
		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{DownloadLink: dbTrack.DownloadLink},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
		})
	}
	fmt.Println(tracks)
	return tracks
}

func allLayoutsToEntity(dbLayouts []db.Layout) []entities.Layout{
	var layouts []entities.Layout
	for _,dbLayout := range dbLayouts{
		layouts = append(layouts, entities.Layout{
			Name:     dbLayout.Name,
			LengthM:  dbLayout.LengthKm,
			Category: entities.TrackCategory{Name: dbLayout.Category},
		})
	}
	return layouts
}

func (t TrackRepositoryImpl) GetTracksByNation(nation string) []entities.Track {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := t.Db.Preload("Layouts").Find(&dbTracks,"nation = ?",nation); result.Error != nil {
		//return error
	}
	for _, dbTrack := range dbTracks {
		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{DownloadLink: dbTrack.DownloadLink},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
		})
	}
	fmt.Println(tracks)
	return tracks
}

func (t TrackRepositoryImpl) GetTracksByLayoutType(category string) []entities.Track {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := t.Db.Preload("Layouts").Joins("join layouts on layouts.track = tracks.name").Where("layouts.category = ?", category).Find(&dbTracks); result.Error != nil {
		//return error
	}
	for _, dbTrack := range dbTracks {

		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{DownloadLink: dbTrack.DownloadLink},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
		})
	}
	fmt.Println(tracks)
	return tracks
}

func (t TrackRepositoryImpl) GetTracksByName(name string) []entities.Track {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := t.Db.Preload("Layouts").Find(&dbTracks,"tracks.name = ?",name); result.Error != nil {
		//return error
	}
	for _, dbTrack := range dbTracks {
		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{DownloadLink: dbTrack.DownloadLink},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
		})
	}
	fmt.Println(tracks)
	return tracks
}

func (t TrackRepositoryImpl) AddNewTrack(track entities.Track) error {
	dbTrack := db.TrackFromEntity(track)
	dbNation := db.NationFromEntity(track.Nation)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := t.Db.Create(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}
