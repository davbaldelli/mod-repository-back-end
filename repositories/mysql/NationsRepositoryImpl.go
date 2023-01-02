package mysql

import (
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/models"
	"gorm.io/gorm"
)

type NationsRepositoryImpl struct {
	Db *gorm.DB
}

func (n NationsRepositoryImpl) SelectAllBrandsNations() ([]models2.Nation, error) {
	var dbNations []models.Nation
	var nations []models2.Nation
	if result := n.Db.Order("nations.name ASC").Distinct("nations.*").Joins("inner join manufacturers on manufacturers.id_nation = nations.id").Find(&dbNations); result.Error != nil {
		return nil, result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, models2.Nation{Name: dbNation.Name, Code: dbNation.Code})
	}
	return nations, nil
}

func (n NationsRepositoryImpl) SelectAllTrackNations() ([]models2.Nation, error) {
	var dbNations []models.Nation
	var nations []models2.Nation
	if result := n.Db.Distinct("nations.*").Joins("inner join tracks on tracks.id_nation = nations.id").Order("nations.name asc").Find(&dbNations); result.Error != nil {
		return nil, result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, models2.Nation{Name: dbNation.Name, Code: dbNation.Code})
	}
	return nations, nil
}
