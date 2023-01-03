package mysql

import (
	"errors"
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math/rand"
)

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (u UserRepositoryImpl) Login(user models2.User) (models2.User, error) {
	var dbUser entities.User
	res := u.Db.Find(&dbUser, "username = ? AND password = SHA2(CONCAT(?, salt),?)", user.Username, user.Password, 224)
	if res.Error != nil {
		return models2.User{}, res.Error
	}
	if res.RowsAffected == 0 {
		return models2.User{}, errors.New("username or password not valid")
	}
	return models2.User{Username: dbUser.Username, Role: models2.Role(dbUser.Role)}, nil
}

func (u UserRepositoryImpl) SignIn(user models2.User) (models2.User, error) {
	salt := randStringRunes(30)
	dbUser := map[string]interface{}{
		"Username": user.Username,
		"Password": clause.Expr{SQL: "SHA2(CONCAT(?, ?),?)", Vars: []interface{}{user.Password, salt, 224}},
		"Role":     string(user.Role),
		"Salt":     salt,
	}
	if res := u.Db.Model(entities.User{}).Create(&dbUser); res.Error != nil {
		return models2.User{}, res.Error
	}
	return models2.User{Username: user.Username, Role: user.Role}, nil
}

func (u UserRepositoryImpl) UpdatePassword(username string, password string) error {
	salt := randStringRunes(30)
	dbUser := map[string]interface{}{
		"Password": clause.Expr{SQL: "SHA2(CONCAT(?, ?),?)", Vars: []interface{}{password, salt, 224}},
		"Salt":     salt,
	}
	if res := u.Db.Model(&entities.User{}).Where("username = ?", username).Updates(&dbUser); res.Error != nil {
		return res.Error
	}
	return nil
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
