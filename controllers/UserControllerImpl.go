package controllers

import (
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type UserControllerImpl struct {
	Repo repositories.UserRepository
}

func (u UserControllerImpl) Login(username string, password string) (models.User, error) {
	return u.Repo.Login(models.User{Username: username, Password: password})
}

func (u UserControllerImpl) SignIn(username string, password string, role models.Role) (models.User, error) {
	return u.Repo.SignIn(models.User{
		Username: username,
		Password: password,
		Role:     role,
	})
}

func (u UserControllerImpl) UpdatePassword(username string, password string) error {
	return u.Repo.UpdatePassword(username, password)
}
