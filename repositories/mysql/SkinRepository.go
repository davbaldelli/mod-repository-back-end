package mysql

import (
	"errors"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type SkinRepositoryImpl struct {
	Db *gorm.DB
}

func (s SkinRepositoryImpl) GetAllSkins() ([]entities.Skin, error) {
	var skins []entities.Skin
	if result := s.Db.Model(&entities.Skin{}).Find(&skins); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return skins, nil
}

func (s SkinRepositoryImpl) SelectCarSkins(carId uint) ([]entities.Skin, error) {
	var skins []entities.Skin
	if result := s.Db.Model(&entities.Skin{}).Where("car_id = ?", carId).Find(&skins); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return skins, nil
}

func (s SkinRepositoryImpl) AddSkin(skin entities.Skin) error {
	if result := s.Db.Create(&skin); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s SkinRepositoryImpl) UpdateSkin(skin entities.Skin) error {
	if result := s.Db.Model(&skin).Where("car_id = ?", skin.CarId).Select("*").Updates(&skin); result.Error != nil {
		return result.Error
	}
	return nil
}
