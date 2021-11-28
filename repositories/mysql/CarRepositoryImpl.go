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

func dbCarToEntity(dbCar db.CarMods) entities.Car {
	return entities.Car{
		Mod: entities.Mod{
			Id: dbCar.Id,
			DownloadLink: dbCar.DownloadLink,
			Premium:      dbCar.Premium,
			Image:        dbCar.Image,
			Author: entities.Author{
				Name: dbCar.Author,
				Link: dbCar.AuthorLink,
			},
		},
		Brand: entities.CarBrand{
			Name:   dbCar.Brand,
			Nation: entities.Nation{Name: dbCar.Nation, Code: dbCar.NationCode},
		},
		ModelName:    dbCar.ModelName,
		Categories:   allCategoriesToEntity(dbCar.Categories),
		Drivetrain:   entities.Drivetrain(dbCar.Drivetrain),
		Transmission: entities.Transmission(dbCar.Transmission),
		Year:         dbCar.Year,
		Torque:       dbCar.Torque,
		TopSpeed:     dbCar.TopSpeed,
		Weight:       dbCar.Weight,
		BHP:          dbCar.BHP,
		CreatedAt: dbCar.CreatedAt,
		UpdatedAt: dbCar.UpdatedAt,
	}
}

func allCategoriesToEntity(dbCategories []db.CarCategory) []entities.CarCategory {
	cats := []entities.CarCategory{}
	for _, dbCat := range dbCategories {
		cats = append(cats, entities.CarCategory{Name: dbCat.Category})
	}
	return cats
}

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
		cars = append(cars, dbCarToEntity(dbCar))
	}
	return cars, nil
}

func (c CarRepositoryImpl) SelectAllCarCategories(premium bool) ([]entities.CarCategory, error) {
	var categories []db.CarCategory
	if result := c.Db.Order("name ASC").Find(&categories); result.Error != nil {
		return nil, result.Error
	}
	return allCategoriesToEntity(categories), nil
}

func (c CarRepositoryImpl) InsertCar(car entities.Car) error {

	dbNation := db.Nation{Name: car.Brand.Nation.Name, Code: car.Brand.Nation.Code}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil{
		return res.Error
	}

	dbBrand := db.Manufacturer{Name: car.Brand.Name, IdNation: dbNation.Id}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbBrand.Name).First(&dbBrand); res.Error != nil{
		return res.Error
	}

	println(dbBrand.Id)

	dbAuthor := db.Author{Name: car.Author.Name, Link: car.Author.Link}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil{
		return res.Error
	}

	dbCar := db.CarFromEntity(car, dbBrand.Id, dbAuthor.Id)

	if res := c.Db.Create(&dbCar); res.Error != nil {
		return res.Error
	}
	return nil
}

func (c CarRepositoryImpl) UpdateCar(car entities.Car) error {
	dbNation := db.Nation{Name: car.Brand.Nation.Name, Code: car.Brand.Nation.Code}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbNation.Name).First(&dbNation); res.Error != nil{
		return res.Error
	}

	dbBrand := db.Manufacturer{Name: car.Brand.Name, IdNation: dbNation.Id}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbBrand.Name).First(&dbBrand); res.Error != nil{
		return res.Error
	}

	dbAuthor := db.Author{Name: car.Author.Name, Link: car.Author.Link}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return res.Error
	}

	if res := c.Db.Where("name = ?", dbAuthor.Name).First(&dbAuthor); res.Error != nil{
		return res.Error
	}

	dbCarUpdated := db.CarFromEntity(car, dbBrand.Id, dbAuthor.Id)


	if res := c.Db.Model(&dbCarUpdated).Select("*").Omit("UpdatedAt", "CreatedAt").Updates(&dbCarUpdated) ; res.Error != nil{
		return res.Error
	}

	if res := c.Db.Model(&dbCarUpdated).Association("Categories").Append(dbCarUpdated.Categories); res != nil{
		return res
	}

	return nil

}

func (c CarRepositoryImpl) SelectAllCars(premium bool) ([]entities.Car, error) {
	return c.selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model) ASC").Preload("Categories")
	}, premium)
}
