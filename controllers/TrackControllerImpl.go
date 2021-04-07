package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetAllTracks() []entities.Track {
	return t.Repo.GetAllTracks()
}

func (t TrackControllerImpl) GetTracksByNation(nationName string) []entities.Track {
	return t.Repo.GetTracksByNation(nationName)
}

func (t TrackControllerImpl) GetTracksByLayoutType(category string) []entities.Track {
	return t.Repo.GetTracksByLayoutType(category)
}

func (t TrackControllerImpl) GetTracksByName(name string) []entities.Track {
	return t.Repo.GetTracksByName(name)
}

func (t TrackControllerImpl) AddTrack(name string, downloadUrl string, layouts []entities.Layout, location string, nation entities.Nation) error {
	return t.Repo.AddNewTrack(entities.Track{
		Mod:      entities.Mod{DownloadLink: downloadUrl},
		Layouts:  layouts,
		Location: location,
		Nation:   nation,
		Name:     name,
	})
}
