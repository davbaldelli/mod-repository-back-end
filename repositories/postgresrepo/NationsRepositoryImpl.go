package postgresrepo

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type NationsRepositoryImpl struct {
	Db *gorm.DB
}

func (n NationsRepositoryImpl) GetAllBrandsNations() []entities.Nation {
	var dbNations []db.Nation
	var nations []entities.Nation
	if result := n.Db.Model(&db.Nation{}).Select("nations.*").Joins("inner join car_brands on car_brands.nation = nations.name").Find(&dbNations); result.Error != nil {
		//return result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations
}

func (n NationsRepositoryImpl) GetAllTrackNations() []entities.Nation {
	var dbNations []db.Nation
	var nations []entities.Nation
	if result := n.Db.Model(&db.Nation{}).Select("nations.*").Joins("inner join tracks on tracks.nation = nations.name").Find(&dbNations); result.Error != nil {
		//return result.Error
	}
	for _, dbNation := range dbNations {
		nations = append(nations, entities.Nation{Name: dbNation.Name})
	}
	return nations
}
