package postgresrepo

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepositoryImpl struct {
	Db *gorm.DB
}

func (c CarRepositoryImpl) AddNewCar(car entities.Car) error {
	dbCar := db.CarFromEntity(car)
	dbNation := db.NationFromEntity(car.Brand.Nation)
	dbBrand := db.BrandFromEntity(car.Brand)

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Create(&dbCar); res.Error != nil {
		return res.Error
	}
	return nil
}

func (c CarRepositoryImpl) GetAllCars() []entities.Car {
	panic("implement me")
}

func (c CarRepositoryImpl) GetCarsByNation(s string) []entities.Car {
	panic("implement me")
}

func (c CarRepositoryImpl) GetCarByModel(s string) []entities.Car {
	panic("implement me")
}

func (c CarRepositoryImpl) GetCarsByBrand(s string) []entities.Car {
	panic("implement me")
}

func (c CarRepositoryImpl) GetCarsByType(s string) []entities.Car {
	panic("implement me")
}
