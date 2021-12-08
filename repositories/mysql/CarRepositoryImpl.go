package mysql

import (
	"errors"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepositoryImpl struct {
	Db *gorm.DB
}

type carsQuery func() *gorm.DB
type selectFromBrandsQuery func(*[]db.Manufacturer) *gorm.DB

func (c CarRepositoryImpl) selectCarsWithQuery(carsQuery carsQuery, premium bool) ([]entities.Car, error) {
	var cars []entities.Car
	var dbCars []db.CarMods

	if premium {
		if result := carsQuery().Find(&dbCars); result.Error != nil {
			return nil, result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	} else {
		if result := carsQuery().Where("premium = false").Find(&dbCars); result.Error != nil {
			return nil, result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	}

	for _, dbCar := range dbCars {
		cars = append(cars, dbCar.ToEntity())
	}
	return cars, nil
}

func (c CarRepositoryImpl) preInsertionQueries(car entities.Car) (db.Car, error) {
	dbNation := db.Nation{Name: car.Brand.Nation.Name, Code: car.Brand.Nation.Code}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return db.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil {
		return db.Car{}, res.Error
	}

	dbBrand := db.Manufacturer{Name: car.Brand.Name, IdNation: dbNation.Id}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return db.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbBrand.Name).First(&dbBrand); res.Error != nil {
		return db.Car{}, res.Error
	}

	println(dbBrand.Id)

	dbAuthor := db.Author{Name: car.Author.Name, Link: car.Author.Link}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return db.Car{}, res.Error
	}

	if res := c.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil {
		return db.Car{}, res.Error
	}

	return db.CarFromEntity(car, dbBrand.Id, dbAuthor.Id), nil
}

func (c CarRepositoryImpl) SelectAllCarCategories() ([]entities.CarCategory, error) {
	return []entities.CarCategory{
		{Name: entities.EnduranceCar},
		{Name: entities.Formula},
		{Name: entities.GT},
		{Name: entities.Touring},
		{Name: entities.Tuned},
		{Name: entities.Vintage},
		{Name: entities.StockCar},
		{Name: entities.Street},
		{Name: entities.RallyCar},
		{Name: entities.Prototype},
	}, nil
}

func (c CarRepositoryImpl) InsertCar(car entities.Car) error {
	if dbCar, err := c.preInsertionQueries(car); err != nil {
		return err
	} else {
		if res := c.Db.Create(&dbCar); res.Error != nil {
			return res.Error
		}
	}
	return nil
}

func (c CarRepositoryImpl) UpdateCar(car entities.Car) (bool, error) {
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

		if res := c.Db.Model(&dbCar).Association("Categories").Append(dbCar.Categories); res != nil {
			return false, res
		}

		return actualCar.Version != dbCar.Version, nil

	}

}

func (c CarRepositoryImpl) SelectAllCars(premium bool) ([]entities.Car, error) {
	return c.selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model) ASC").Preload("Categories")
	}, premium)
}
