package controllers

import (
	"github.com/davide/ModRepository/models/entities"
	"github.com/davide/ModRepository/repositories"
)

type UserControllerImpl struct {
	Repo repositories.UserRepository
}

func (u UserControllerImpl) Login(username string, password string) (entities.User, error) {
	return u.Repo.Login(entities.User{Username: username, Password: password})
}

func (u UserControllerImpl) SignIn(username string, password string, role entities.Role) (entities.User, error)  {
	return u.Repo.SignIn(entities.User{
		Username: username,
		Password: password,
		Role:     role,
	})
}

