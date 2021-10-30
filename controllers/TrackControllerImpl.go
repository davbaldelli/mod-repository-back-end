package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetTrackByName(name string) (entities.Track, error) {
	return t.Repo.SelectTrackByName(name)
}

func (t TrackControllerImpl) GetTracksByTag(tag entities.TrackTag, premium bool) ([]entities.Track, error) {
	return t.Repo.SelectTracksByTag(tag, premium)
}

func (t TrackControllerImpl) GetAllTracks(premium bool) ([]entities.Track, error) {
	return t.Repo.SelectAllTracks(premium)
}

func (t TrackControllerImpl) GetTracksByNation(nationName string,premium bool) ([]entities.Track, error) {
	return t.Repo.SelectTracksByNation(nationName, premium)
}

func (t TrackControllerImpl) GetTracksByLayoutType(category string,premium bool) ([]entities.Track, error) {
	return t.Repo.SelectTracksByLayoutType(category, premium)
}

func (t TrackControllerImpl) GetTracksByName(name string, premium bool) ([]entities.Track, error) {
	return t.Repo.SelectTracksByName(name, premium)
}

func (t TrackControllerImpl) AddTrack(track entities.Track) error {
	return t.Repo.InsertTrack(track)
}
