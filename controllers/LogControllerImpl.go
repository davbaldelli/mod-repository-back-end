package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type LogControllerImpl struct {
	Repo repositories.LogRepository
}

func (l LogControllerImpl) GetTrackLogs(b bool) ([]entities.TrackLog, error) {
	return l.Repo.SelectAllTrackLogs(b)
}

func (l LogControllerImpl) GetCarLogs(b bool) ([]entities.CarLog, error) {
	return l.Repo.SelectAllCarLogs(b)
}
