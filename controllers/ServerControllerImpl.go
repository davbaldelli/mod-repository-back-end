package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type ServersControllerImpl struct {
	Repo repositories.ServersRepository
}

func (s ServersControllerImpl) AddServer(server entities.Server) error {
	return s.Repo.AddServer(server)
}

func (s ServersControllerImpl) UpdateServer(server entities.Server) error {
	return s.Repo.UpdateServer(server)
}

func (s ServersControllerImpl) GetAllServers() ([]entities.Server, error) {
	return s.Repo.GetAllServers()
}
