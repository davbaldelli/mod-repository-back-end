package postgresrepo

import (
	"errors"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) Login(user entities.User) (entities.User, error) {
	var dbUser db.User
	res := u.Db.Find(&dbUser, "username = ? AND password = crypt(?,password)",user.Username, user.Password)
	if  res.Error != nil{
		return entities.User{},res.Error
	}
	if res.RowsAffected == 0 {
		return entities.User{}, errors.New("username or password not valid")
	}
	return entities.User{Username: dbUser.Username, IsAdmin: dbUser.IsAdmin}, nil
}

