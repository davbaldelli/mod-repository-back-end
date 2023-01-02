package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type ServersControllerImpl struct {
	Repo repositories.ServersRepository
}

func (s ServersControllerImpl) AddServer(server models.Server) error {
	return s.Repo.AddServer(server)
}

func (s ServersControllerImpl) UpdateServer(server models.Server) error {
	return s.Repo.UpdateServer(server)
}

func (s ServersControllerImpl) GetAllServers() ([]models.Server, error) {
	return s.Repo.GetAllServers()
}

func (s ServersControllerImpl) DeleteServer(server models.Server) error {
	return s.Repo.DeleteServer(server)
}
