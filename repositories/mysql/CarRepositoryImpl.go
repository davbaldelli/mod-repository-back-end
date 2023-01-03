package mysql

import (
	"errors"
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepositoryImpl struct {
	Db *gorm.DB
}

type carsQuery func() *gorm.DB
type selectFromBrandsQuery func(*[]entities.Manufacturer) *gorm.DB

func (c CarRepositoryImpl) selectCarsWithQuery(carsQuery carsQuery, premium bool, admin bool) ([]models2.Car, error) {
	var cars []models2.Car
	var dbCars []entities.CarMods

	if result := carsQuery().Find(&dbCars); result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected == 0 {
		return nil, errors.New("not found")
	}

	for _, dbCar := range dbCars {
		cars = append(cars, dbCar.ToEntity(premium, admin))
	}
	return cars, nil
}

func (c CarRepositoryImpl) preInsertionQueries(car models2.Car) (entities.Car, error) {
	dbNation := entities.NationFromEntity(car.Brand.Nation)

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return entities.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil {
		return entities.Car{}, res.Error
	}

	dbBrand := entities.ManufacturerFromEntity(car.Brand, dbNation.Id)

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return entities.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbBrand.Name).First(&dbBrand); res.Error != nil {
		return entities.Car{}, res.Error
	}

	println(dbBrand.Id)

	dbAuthor := entities.AuthorFromEntity(car.Author)

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return entities.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil {
		return entities.Car{}, res.Error
	}

	return entities.CarFromEntity(car, dbBrand.Id, dbAuthor.Id), nil
}

func (c CarRepositoryImpl) SelectAllCarCategories() ([]models2.CarCategory, error) {
	return []models2.CarCategory{
		{Name: models2.EnduranceCar},
		{Name: models2.OpenWheel},
		{Name: models2.GT},
		{Name: models2.Touring},
		{Name: models2.Tuned},
		{Name: models2.Vintage},
		{Name: models2.StockCar},
		{Name: models2.Street},
		{Name: models2.RallyCar},
		{Name: models2.Prototype},
	}, nil
}

func (c CarRepositoryImpl) InsertCar(car *models2.Car) error {
	if dbCar, err := c.preInsertionQueries(*car); err != nil {
		return err
	} else {
		if res := c.Db.Create(&dbCar); res.Error != nil {
			return res.Error
		}
		car.Id = dbCar.Id
	}
	return nil
}

func (c CarRepositoryImpl) UpdateCar(car models2.Car) (bool, error) {
	if dbCar, err := c.preInsertionQueries(car); err != nil {
		return false, err
	} else {

		actualCar := dbCar

		if res := c.Db.First(&actualCar, car.Id); res.Error != nil {
			return false, res.Error
		}

		if res := c.Db.Model(&dbCar).Clauses(clause.Returning{Columns: []clause.Column{{Name: "version"}}}).Select("*").Omit("UpdatedAt", "CreatedAt").Updates(&dbCar); res.Error != nil {
			return false, res.Error
		}

		if res := c.Db.Where("car_id = ?", dbCar.Id).Delete(&entities.CarCategory{}); res.Error != nil {
			return false, res.Error
		}

		if res := c.Db.Model(&dbCar).Association("Categories").Append(dbCar.Categories); res != nil {
			return false, res
		}

		if res := c.Db.Where("car_id = ?", dbCar.Id).Delete(&entities.CarImage{}); res.Error != nil {
			return false, res.Error
		}

		if res := c.Db.Model(&dbCar).Association("Images").Append(dbCar.Images); res != nil {
			return false, res
		}

		return actualCar.Version != dbCar.Version, nil

	}

}

func (c CarRepositoryImpl) SelectAllCars(premium bool, admin bool) ([]models2.Car, error) {
	return c.selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model) ASC").Preload("Categories").Preload("Images")
	}, premium, admin)
}
