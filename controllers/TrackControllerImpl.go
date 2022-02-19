package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetAllTracks(role entities.Role) ([]entities.Track, error) {
	return t.Repo.SelectAllTracks(role)
}

func (t TrackControllerImpl) AddTrack(track *entities.Track) error {
	return t.Repo.InsertTrack(track)
}

func (t TrackControllerImpl) UpdateTrack(track entities.Track) (bool, error) {
	return t.Repo.UpdateTrack(track)
}
