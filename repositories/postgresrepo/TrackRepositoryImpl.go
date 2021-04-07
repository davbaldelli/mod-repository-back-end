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

func (t TrackRepositoryImpl) GetAllTracks() []entities.Track {
	panic("implement me")
}

func (t TrackRepositoryImpl) GetTracksByNation(s string) []entities.Track {
	panic("implement me")
}

func (t TrackRepositoryImpl) GetTracksByLayoutType(s string) []entities.Track {
	panic("implement me")
}

func (t TrackRepositoryImpl) GetTracksByName(s string) []entities.Track {
	panic("implement me")
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
