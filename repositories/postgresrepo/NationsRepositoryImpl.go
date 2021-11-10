package postgresrepo

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type NationsRepositoryImpl struct {
	Db *gorm.DB
}

func (n NationsRepositoryImpl) SelectAllBrandsNations() ([]entities.Nation, error) {
	var dbNations []db.Nation
	var nations []entities.Nation
	if result := n.Db.Order("nations.name ASC").Distinct("nations.name").Joins("inner join manufacturers on manufacturers.id_nation = nations.id").Find(&dbNations); result.Error != nil {
		return nil, result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations, nil
}

func (n NationsRepositoryImpl) SelectAllTrackNations() ([]entities.Nation, error) {
	var dbNations []db.Nation
	var nations []entities.Nation
	if result := n.Db.Distinct("nations.name").Joins("inner join tracks on tracks.id_nation = nations.id").Order("nations.name asc").Find(&dbNations); result.Error != nil {
		return nil, result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations, nil
}
