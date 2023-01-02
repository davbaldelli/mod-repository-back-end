package mysql

import (
	"errors"
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TrackRepositoryImpl struct {
	Db *gorm.DB
}

type selectFromTrackQuery func() *gorm.DB

func selectTracksWithQuery(query selectFromTrackQuery, premium bool, admin bool) ([]models2.Track, error) {
	var dbTracks []models.TrackMod
	var tracks []models2.Track

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

func (t TrackRepositoryImpl) preInsertionQueries(track models2.Track) (models.Track, error) {
	dbNation := models.NationFromEntity(track.Nation)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return models.Track{}, res.Error
	}

	if res := t.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil {
		return models.Track{}, res.Error
	}

	dbAuthor := models.AuthorFromEntity(track.Author)

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return models.Track{}, res.Error
	}

	if res := t.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil {
		return models.Track{}, res.Error
	}

	return models.TrackFromEntity(track, dbNation.Id, dbAuthor.Id), nil
}

func (t TrackRepositoryImpl) SelectAllTracks(premium bool, admin bool) ([]models2.Track, error) {
	return selectTracksWithQuery(func() *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Preload("Tags").Preload("Images")
	}, premium, admin)
}

func (t TrackRepositoryImpl) InsertTrack(track *models2.Track) error {

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

func (t TrackRepositoryImpl) UpdateTrack(track models2.Track) (bool, error) {

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

		if res := t.Db.Where("id_track = ?", dbTrack.Id).Delete(&models.Layout{}); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Association("Layouts").Append(dbTrack.Layouts); res != nil {
			return false, res
		}

		if res := t.Db.Where("id_track = ?", dbTrack.Id).Delete(&models.TrackTag{}); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Association("Tags").Append(dbTrack.Tags); res != nil {
			return false, res
		}

		if res := t.Db.Where("track_id = ?", dbTrack.Id).Delete(&models.TrackImage{}); res.Error != nil {
			return false, res.Error
		}

		if res := t.Db.Model(&dbTrack).Association("Images").Append(dbTrack.Images); res != nil {
			return false, res
		}

		return oldTrack.Version != track.Version, nil
	}

}
