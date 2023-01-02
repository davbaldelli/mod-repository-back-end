package controllers

import (
	"github.com/davide/ModRepository/controllers/helpers"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	Repo repositories.TrackRepository
}

func (t TrackControllerImpl) GetAllTracks(role models.Role) ([]models.Track, error) {
	return t.Repo.SelectAllTracks(helpers.IsPremium(role), helpers.IsAdmin(role))
}

func (t TrackControllerImpl) AddTrack(track *models.Track) error {
	return t.Repo.InsertTrack(track)
}

func (t TrackControllerImpl) UpdateTrack(track models.Track) (bool, error) {
	return t.Repo.UpdateTrack(track)
}
