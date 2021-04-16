package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}


func (t TrackControllerImpl) GetAllTracks() ([]entities.Track, error) {
	return t.Repo.SelectAllTracks()
}

func (t TrackControllerImpl) GetTracksByNation(nationName string) ([]entities.Track, error) {
	return t.Repo.SelectTracksByNation(nationName)
}

func (t TrackControllerImpl) GetTracksByLayoutType(category string) ([]entities.Track, error) {
	return t.Repo.SelectTracksByLayoutType(category)
}

func (t TrackControllerImpl) GetTracksByName(name string) ([]entities.Track, error) {
	return t.Repo.SelectTracksByName(name)
}

func (t TrackControllerImpl) AddTrack(track entities.Track) error {
	return t.Repo.InsertTrack(track)
}
