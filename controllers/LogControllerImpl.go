package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type LogControllerImpl struct {
	Repo repositories.LogRepository
}

func (l LogControllerImpl) GetTrackLogs() ([]entities.TrackLog, error) {
	return l.Repo.SelectAllTrackLogs()
}

func (l LogControllerImpl) GetCarLogs() ([]entities.CarLog, error) {
	return l.Repo.SelectAllCarLogs()
}
