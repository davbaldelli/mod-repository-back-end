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
type selectFromAuthorsQuery func(*[]db.Author, []string) *gorm.DB



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
	}, func(authors *[]db.Author, tracksNames []string) *gorm.DB {
		return t.Db.Joins("join tracks on author = authors.name").Find(authors,"tracks.name IN ?",tracksNames)
	})
}

func (t TrackRepositoryImpl) SelectTracksByNation(nation string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"nation = ?",nation)
	}, func(authors *[]db.Author, tracksNames []string) *gorm.DB {
		return t.Db.Joins("join tracks on author = authors.name").Find(authors,"tracks.name IN ?",tracksNames)
	})
}

func (t TrackRepositoryImpl) SelectTracksByLayoutType(category string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Distinct().Preload("Layouts").Joins("join layouts on layouts.track = tracks.name").Where("layouts.category = ?", category).Find(&tracks)
	}, func(authors *[]db.Author, tracksNames []string) *gorm.DB {
		return t.Db.Joins("join tracks on author = authors.name").Find(authors,"tracks.name IN ?",tracksNames)
	})
}

func (t TrackRepositoryImpl) SelectTracksByName(name string) ([]entities.Track,error) {
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks,"LOWER(tracks.name) LIKE LOWER(?)","%"+name+"%")
	}, func(authors *[]db.Author, tracksNames []string) *gorm.DB {
		return t.Db.Joins("join tracks on author = authors.name").Find(authors,"tracks.name IN ?",tracksNames)
	})
}

func (t TrackRepositoryImpl) SelectTrackByTag(tag entities.TrackTag) ([]entities.Track,error){
	return selectTracksWithQuery(func(tracks *[]db.Track) *gorm.DB {
		return t.Db.Order("name ASC").Preload("Layouts").Find(&tracks," ? = ANY (tags)", tag)
	}, func(authors *[]db.Author, tracksNames []string) *gorm.DB {
		return t.Db.Joins("join tracks on author = authors.name").Find(authors,"tracks.name IN ?",tracksNames)
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

func selectTracksWithQuery(query selectFromTrackQuery, authorsQuery selectFromAuthorsQuery) ([]entities.Track, error) {
	var dbTracks []db.Track
	var tracks []entities.Track
	if result := query(&dbTracks); result.Error != nil {
		return nil,result.Error
	}

	var tracksNames []string
	for _,dbTrack := range dbTracks{
		tracksNames = append(tracksNames, dbTrack.Name)
	}
	var dbAuthors []db.Author
	if result := authorsQuery(&dbAuthors,tracksNames); result.Error != nil{
		return nil,result.Error
	}

	authorMap := make(map[string]db.Author)

	for _, author := range dbAuthors {
		authorMap[author.Name] = author
	}

	for _, dbTrack := range dbTracks {
		var tags []entities.TrackTag
		for _, tag := range dbTrack.Tags {
			tags = append(tags, entities.TrackTag(tag))
		}
		tracks = append(tracks, entities.Track{
			Mod:      entities.Mod{
				DownloadLink: dbTrack.DownloadLink,
				Premium: dbTrack.Premium,
				Image: dbTrack.Image,
				Author: entities.Author{
					Name: authorMap[dbTrack.Author].Name,
					Link: authorMap[dbTrack.Author].Link,
				},
			},
			Name:     dbTrack.Name,
			Layouts:  allLayoutsToEntity(dbTrack.Layouts),
			Location: dbTrack.Location,
			Nation:   entities.Nation{Name: dbTrack.Nation},
			Tags: tags,
		})
	}
	return tracks,nil
}
