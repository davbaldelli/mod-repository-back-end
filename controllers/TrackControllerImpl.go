package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetAllTracks() []models.Track {
	return t.Repo.GetAllTracks()
}

func (t TrackControllerImpl) GetTracksByNation(nationName string) []models.Track {
	return t.Repo.GetTracksByNation(nationName)
}

func (t TrackControllerImpl) GetTracksByLayoutType(category string) []models.Track {
	return t.Repo.GetTracksByLayoutType(category)
}

func (t TrackControllerImpl) GetTracksByName(name string) []models.Track {
	return t.Repo.GetTracksByName(name)
}

func (t TrackControllerImpl) AddTrack(name string, downloadUrl string, layouts []models.Layout, location models.Location) error {
	return t.Repo.AddNewTrack(models.Track{
		Mod:      models.Mod{DownloadLink: downloadUrl},
		Layouts:  layouts,
		Location: location,
		Name:     name,
	})
}
