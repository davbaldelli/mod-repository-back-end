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
	if result := n.Db.Distinct("nations.name").Joins("inner join car_brands on car_brands.nation = nations.name").Find(&dbNations); result.Error != nil {
		return nil,result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations,nil
}

func (n NationsRepositoryImpl) SelectAllTrackNations() ([]entities.Nation, error) {
	var dbNations []db.Nation
	var nations []entities.Nation
	if result := n.Db.Distinct("nations.name").Joins("inner join tracks on tracks.nation = nations.name").Find(&dbNations); result.Error != nil {
		return nil,result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations,nil
}
