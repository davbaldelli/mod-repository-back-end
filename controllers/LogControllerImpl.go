package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type LogControllerImpl struct {
	Repo repositories.LogRepository
}

func (l LogControllerImpl) GetTrackLogs() ([]models.TrackLog, error) {
	return l.Repo.SelectAllTrackLogs()
}

func (l LogControllerImpl) GetCarLogs() ([]models.CarLog, error) {
	return l.Repo.SelectAllCarLogs()
}
