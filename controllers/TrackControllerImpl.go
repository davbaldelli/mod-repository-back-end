package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetAllTracks() []models.Track {
	return t.repo.GetAllTracks()
}

func (t TrackControllerImpl) GetTracksByNation(nationName string) []models.Track {
	return t.repo.GetTracksByNation(nationName)
}

func (t TrackControllerImpl) GetTracksByLayoutType(category string) []models.Track {
	return t.repo.GetTracksByLayoutType(category)
}

func (t TrackControllerImpl) GetTracksByName(name string) []models.Track {
	return t.repo.GetTracksByName(name)
}

func (t TrackControllerImpl) AddTrack(name string, downloadUrl string, layouts []models.Layout, location models.Location) error {
	return t.repo.AddNewTrack(models.Track{
		Mod:      models.Mod{DownloadLink: downloadUrl},
		Layouts:  layouts,
		Location: location,
		Name:     name,
	})
}
