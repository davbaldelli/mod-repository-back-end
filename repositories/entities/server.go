package entities

import (
	"github.com/davide/ModRepository/models"
)

type Server struct {
	Id               uint `gorm:"primaryKey"`
	Name             string
	Description      string
	JoinLink         string
	Password         string
	Online           bool
	TrackId          uint
	OutsideTrackName string        `gorm:"column:outside_track_name"`
	OutsideTrackLink string        `gorm:"column:outside_track_link"`
	OutsideTrack     bool          `gorm:"column:outside_track"`
	Cars             []*Car        `gorm:"many2many:server_cars;foreignKey:Id;joinForeignKey:ServerId;References:Id;joinReferences:CarId"`
	OutsideCars      []*OutsideMod `gorm:"foreignKey:ServerId"`
}

func ServerFromEntity(server models.Server) Server {
	return Server{
		Id:               server.Id,
		Name:             server.Name,
		Description:      server.Description,
		JoinLink:         server.JoinLink,
		Password:         server.Password,
		Online:           server.Online,
		TrackId:          server.Track,
		OutsideTrack:     server.OutsideTrack,
		OutsideTrackName: server.OutsideTrackName,
		OutsideTrackLink: server.OutsideTrackLink,
	}
}

type OutsideMod struct {
	Id           string `gorm:"primaryKey"`
	Name         string
	DownloadLink string
	ServerId     uint
}

func (s Server) ToEntity() models.Server {
	var cars []uint
	cars = make([]uint, 0)
	var outsideCars []models.OutsideMod
	outsideCars = make([]models.OutsideMod, 0)
	for _, dbCar := range s.Cars {
		cars = append(cars, dbCar.Id)
	}
	for _, outsideCar := range s.OutsideCars {
		outsideCars = append(outsideCars, outsideCar.ToEntity())
	}
	return models.Server{
		Id:               s.Id,
		Name:             s.Name,
		Description:      s.Description,
		JoinLink:         s.JoinLink,
		Password:         s.Password,
		Online:           s.Online,
		Track:            s.TrackId,
		OutsideTrack:     s.OutsideTrack,
		OutsideTrackName: s.OutsideTrackName,
		OutsideTrackLink: s.OutsideTrackLink,
		Cars:             cars,
		OutsideCars:      outsideCars,
	}
}

func OutsideModFromEntity(mod models.OutsideMod, serverId uint) OutsideMod {
	return OutsideMod{
		Name:         mod.Name,
		DownloadLink: mod.DownloadLink,
		ServerId:     serverId,
	}
}

func (o OutsideMod) ToEntity() models.OutsideMod {
	return models.OutsideMod{
		Id:           o.Id,
		Name:         o.Name,
		DownloadLink: o.DownloadLink,
		ServerId:     o.ServerId,
	}
}
