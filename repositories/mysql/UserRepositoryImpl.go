package mysql

import (
	"errors"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) Login(user entities.User) (entities.User, error) {
	var dbUser db.User
	res := u.Db.Find(&dbUser, "username = ? AND password = SHA2(CONCAT(?, salt),?)", user.Username, user.Password, 224)
	if res.Error != nil {
		return entities.User{}, res.Error
	}
	if res.RowsAffected == 0 {
		return entities.User{}, errors.New("username or password not valid")
	}
	return entities.User{Username: dbUser.Username, Role: entities.Role(dbUser.Role)}, nil
}

func (u UserRepositoryImpl) SignIn(user entities.User) (entities.User, error) {
	salt := randStringRunes(30)
	dbUser := map[string]interface{}{
		"Username": user.Username,
		"Password": clause.Expr{SQL: "SHA2(CONCAT(?, ?),?)", Vars: []interface{}{user.Password, salt, 224}},
		"Role":     string(user.Role),
		"Salt":     salt,
	}
	if res := u.Db.Model(db.User{}).Create(&dbUser); res.Error != nil {
		return entities.User{}, res.Error
	}
	return entities.User{Username: user.Username, Role: user.Role}, nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
