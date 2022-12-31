package models

import "github.com/davide/ModRepository/models/entities"

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

func ServerFromEntity(server entities.Server) Server {
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

func (s Server) ToEntity() entities.Server {
	var cars []uint
	cars = make([]uint, 0)
	var outsideCars []entities.OutsideMod
	outsideCars = make([]entities.OutsideMod, 0)
	for _, dbCar := range s.Cars {
		cars = append(cars, dbCar.Id)
	}
	for _, outsideCar := range s.OutsideCars {
		outsideCars = append(outsideCars, outsideCar.ToEntity())
	}
	return entities.Server{
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

func OutsideModFromEntity(mod entities.OutsideMod, serverId uint) OutsideMod {
	return OutsideMod{
		Name:         mod.Name,
		DownloadLink: mod.DownloadLink,
		ServerId:     serverId,
	}
}

func (o OutsideMod) ToEntity() entities.OutsideMod {
	return entities.OutsideMod{
		Id:           o.Id,
		Name:         o.Name,
		DownloadLink: o.DownloadLink,
		ServerId:     o.ServerId,
	}
}
