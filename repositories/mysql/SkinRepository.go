package mysql

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"gorm.io/gorm"
)

type SkinRepositoryImpl struct {
	Db *gorm.DB
}

func (s SkinRepositoryImpl) GetAllSkins() ([]models.Skin, error) {
	var skins []models.Skin
	if result := s.Db.Model(&models.Skin{}).Find(&skins); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return skins, nil
}

func (s SkinRepositoryImpl) SelectCarSkins(carId uint) ([]models.Skin, error) {
	var skins []models.Skin
	if result := s.Db.Model(&models.Skin{}).Where("car_id = ?", carId).Find(&skins); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}
	return skins, nil
}

func (s SkinRepositoryImpl) AddSkin(skin models.Skin) error {
	if result := s.Db.Create(&skin); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s SkinRepositoryImpl) UpdateSkin(skin models.Skin) error {
	if result := s.Db.Model(&skin).Where("car_id = ?", skin.CarId).Select("*").Updates(&skin); result.Error != nil {
		return result.Error
	}
	return nil
}
