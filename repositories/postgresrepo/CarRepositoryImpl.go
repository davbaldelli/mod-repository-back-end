package postgresrepo

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

func dbCarToEntity(dbCar db.CarMods, categories []db.CarCategory)entities.Car{
	return entities.Car{
		Mod: entities.Mod{
			DownloadLink: dbCar.DownloadLink,
			Premium: dbCar.Premium,
			Image: dbCar.Image,
			Author: entities.Author{
				Name: dbCar.Author,
				Link: dbCar.AuthorLink,
			},
		},
		Brand: entities.CarBrand{
			Name:   dbCar.Brand,
			Nation: entities.Nation{Name: dbCar.Nation},
		},
		ModelName:  dbCar.ModelName,
		Categories: allCategoriesToEntity(categories),
		Drivetrain: entities.Drivetrain(dbCar.Drivetrain),
		Transmission: entities.Transmission(dbCar.Transmission),
		Year: dbCar.Year,
		Torque: dbCar.Torque,
		TopSpeed: dbCar.TopSpeed,
		Weight: dbCar.Weight,
		BHP: dbCar.BHP,

	}
}

func allCategoriesToEntity(dbCategories []db.CarCategory) []entities.CarCategory{
	var cats []entities.CarCategory
	for _,dbCat := range  dbCategories {
		cats = append(cats, entities.CarCategory{Name: dbCat.Name})
	}
	return cats
}


func (c CarRepositoryImpl)selectCarsWithQuery(carsQuery carsQuery, premium bool) ([]entities.Car, error){
	var cars []entities.Car
	var dbCars []db.CarMods

	if premium {
		if result := carsQuery().Find(&dbCars); result.Error != nil{
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	} else {
		if result := carsQuery().Where("premium = false").Find(&dbCars); result.Error != nil{
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	}

	for _, dbCar := range dbCars {
		var categories []db.CarCategory
		var catId []uint
		for _, catAss := range dbCar.Categories{
			catId = append(catId, catAss.IdCategory)
		}
		if res := c.Db.Where("id IN (?)", catId).Find(&categories); res.Error != nil{
			return nil, res.Error
		}
		cars = append(cars, dbCarToEntity(dbCar, categories))
	}
	return cars,nil
}

func (c CarRepositoryImpl) SelectAllCarCategories(premium bool) ([]entities.CarCategory, error) {
	var categories []db.CarCategory
	if result := c.Db.Order("name ASC").Find(&categories) ; result.Error != nil{
		return  nil, result.Error
	}
	return allCategoriesToEntity(categories), nil
}

func (c CarRepositoryImpl) InsertCar(car entities.Car) error {

	dbNation := db.Nation{Name: car.Brand.Nation.Name}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbNation); res.Error != nil {
		return res.Error
	}

	dbBrand := db.Manufacturer{Name: car.Brand.Name, IdNation: dbNation.Id}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbBrand); res.Error != nil {
		return res.Error
	}

	dbAuthor := db.Author{ Name:   car.Author.Name, Link:   car.Author.Link}

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&dbAuthor); res.Error != nil {
		return res.Error
	}

	dbCar := db.CarFromEntity(car, dbBrand.Id, dbAuthor.Id)

	if res := c.Db.Create(&dbCar); res.Error != nil {
		return res.Error
	}
	return nil
}

func (c CarRepositoryImpl) SelectAllCars(premium bool) ([]entities.Car,error) {
	return c.selectCarsWithQuery(func() *gorm.DB {
			return c.Db.Order("concat(brand,' ',model) ASC").Preload(clause.Associations)
	}, premium)
}
