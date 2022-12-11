package mysql

import (
	"errors"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrackRepositoryImpl struct {
	Db *gorm.DB
}

type selectFromTrackQuery func() *gorm.DB

func selectTracksWithQuery(query selectFromTrackQuery, premium bool, admin bool) ([]entities.Track, error) {
	var dbTracks []db.TrackMod
	var tracks []entities.Track

	if result := query().Find(&dbTracks); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	for _, dbTrack := range dbTracks {
		tracks = append(tracks, dbTrack.ToEntity(premium, admin))
	}
	return tracks, nil
}

func (t TrackRepositoryImpl) preInsertionQueries(track entities.Track) (db.Track, error) {
	dbNation := db.NationFromEntity(track.Nation)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return db.Track{}, res.Error
	}

	if res := t.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil {
		return db.Track{}, res.Error
	}

	dbAuthor := db.AuthorFromEntity(track.Author)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return db.Track{}, res.Error
	}

	if res := t.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil {
		return db.Track{}, res.Error
	}

	return db.TrackFromEntity(track, dbNation.Id, dbAuthor.Id), nil
}

func (t TrackRepositoryImpl) SelectAllTracks(premium bool, admin bool) ([]entities.Track, error) {
	return selectTracksWithQuery(func() *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Preload("Tags")
	}, premium, admin)
}

func (t TrackRepositoryImpl) InsertTrack(track *entities.Track) error {

	if dbTrack, err := t.preInsertionQueries(*track); err != nil {
		return err
	} else {
		if res := t.Db.Create(&dbTrack); res.Error != nil {
			return res.Error
		}
		track.Id = dbTrack.Id
	}
	return nil
}

func (t TrackRepositoryImpl) UpdateTrack(track entities.Track) (bool, error) {

	if dbTrack, err := t.preInsertionQueries(track); err != nil {
		return false, err
	} else {
		oldTrack := dbTrack

		if res := t.Db.First(&oldTrack, track.Id); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Select("*").Omit("UpdatedAt", "CreatedAt").Updates(&dbTrack); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Where("id_track = ?", dbTrack.Id).Delete(&db.Layout{}); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Association("Layouts").Append(dbTrack.Layouts); res != nil {
			return false, res
		}

		if res := t.Db.Where("id_track = ?", dbTrack.Id).Delete(&db.TrackTag{}); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Association("Tags").Append(dbTrack.Tags); res != nil {
			return false, res
		}

		return oldTrack.Version != track.Version, nil
	}

}
