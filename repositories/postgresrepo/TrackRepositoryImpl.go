package postgresrepo

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

func trackToEntity(dbTrack db.TrackMod) entities.Track {
	var tags []entities.TrackTag
	/*for _, tag := range dbTrack.Tags {
		tags = append(tags, entities.TrackTag(tag))
	}*/
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
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Track{}, errors.New("not found")
		}
		return entities.Track{}, result.Error
	}

	return trackToEntity(track), nil
}

func (t TrackRepositoryImpl) SelectAllTracks(premium bool) ([]entities.Track,error) {
	return selectTracksWithQuery(func() *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts")
	}, premium)
}


func (t TrackRepositoryImpl) InsertTrack(track entities.Track) error {

	dbNation := db.Nation{Name: track.Nation.Name}

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {return res.Error}

	dbAuthor := db.Author{Name: track.Author.Name, Link: track.Author.Link}

	if res := t.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return res.Error
	}

	dbTrack := db.TrackFromEntity(track, dbNation.Id, dbAuthor.Id)

	if res := t.Db.Create(&dbTrack); res.Error != nil {
		return res.Error
	}
	return nil
}

func selectTracksWithQuery(query selectFromTrackQuery, premium bool) ([]entities.Track, error) {
	var dbTracks []db.TrackMod
	var tracks []entities.Track
	if premium {
		if result := query().Find(&dbTracks); result.Error != nil {
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	} else {
		if result := query().Where("premium = false").Find(&dbTracks); result.Error != nil {
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	}
	for _, dbTrack := range dbTracks {
		tracks = append(tracks, trackToEntity(dbTrack))
	}
	return tracks,nil
}
