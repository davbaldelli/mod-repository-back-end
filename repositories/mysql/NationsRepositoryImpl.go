package mysql

import (
	"github.com/davide/ModRepository/models"
	"gorm.io/gorm"
)

type NationsRepositoryImpl struct {
	Db *gorm.DB
}

func (n NationsRepositoryImpl) SelectAllBrandsNations() ([]models.Nation, error) {
	var nations []models.Nation
	if result := n.Db.Order("nations.name ASC").Distinct("nations.*").Joins("inner join manufacturers on manufacturers.id_nation = nations.id").Find(&nations); result.Error != nil {
		return nil, result.Error
	}
	return nations, nil
}

func (n NationsRepositoryImpl) SelectAllTrackNations() ([]models.Nation, error) {
	var nations []models.Nation
	if result := n.Db.Distinct("nations.*").Joins("inner join tracks on tracks.id_nation = nations.id").Order("nations.name asc").Find(&nations); result.Error != nil {
		return nil, result.Error
	}
	return nations, nil
}
