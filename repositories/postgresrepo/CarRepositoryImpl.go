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
type selectFromBrandsQuery func(*[]db.CarBrand) *gorm.DB

func dbCarToEntity(dbCar db.CarMods)entities.Car{
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
		Categories: allCategoriesToEntity(dbCar.Categories),
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


func selectCarsWithQuery(carsQuery carsQuery, premium bool) ([]entities.Car, error){
	var cars []entities.Car
	var dbCars []db.CarMods


	if premium {
		if result := carsQuery().Find(&dbCars); result.Error != nil{
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	} else {
		if result := carsQuery().Where("cars.premium = false").Find(&dbCars); result.Error != nil{
			return nil,result.Error
		} else if result.RowsAffected == 0 {
			return nil, errors.New("not found")
		}
	}

	for _, dbCar := range dbCars {
		cars = append(cars, dbCarToEntity(dbCar))
	}
	return cars,nil
}

func (c CarRepositoryImpl) SelectCarByModel(model string) (entities.Car, error) {
	dbCar := db.CarMods{ModelName: model}
	if result := c.Db.Preload("Categories").First(&dbCar); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Car{}, errors.New("not found")
		}
		return entities.Car{}, result.Error
	}

	return dbCarToEntity(dbCar), nil
}

func (c CarRepositoryImpl) SelectAllCarCategories() ([]entities.CarCategory, error) {
	var categories []db.CarCategory
	if result := c.Db.Order("name ASC").Find(&categories) ; result.Error != nil{
		return  nil, result.Error
	}
	return allCategoriesToEntity(categories), nil
}

func (c CarRepositoryImpl) InsertCar(car entities.Car) error {
	dbCar := db.CarFromEntity(car)
	dbNation := db.NationFromEntity(car.Brand.Nation)
	dbBrand := db.BrandFromEntity(car.Brand)

	if res := c.Db.Clauses(clause.OnConflict{DoNothing: true}).Create(&car.Author); res.Error != nil {
		return res.Error
	}

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

func (c CarRepositoryImpl) SelectAllCars(premium bool) ([]entities.Car,error) {
	return selectCarsWithQuery(func() *gorm.DB {
			return c.Db.Order("concat(brand,' ',model_name) ASC").Preload("Categories")
	}, premium)
}


func (c CarRepositoryImpl) SelectCarsByNation(nation string, premium bool) ([]entities.Car,error) {
	return selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model_name) ASC").Preload("Categories").Where("nation = ?",nation)
	},premium)

}

func (c CarRepositoryImpl) SelectCarsByModelName(model string, premium bool) ([]entities.Car,error) {
	return selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model_name) ASC").Preload("Categories").Where("LOWER(concat(brand,' ',model_name)) LIKE LOWER(?)", "%"+model+"%")
	}, premium)
}

func (c CarRepositoryImpl) SelectCarsByBrand(brandName string, premium bool) ([]entities.Car,error) {
	return selectCarsWithQuery(func() *gorm.DB {
		return c.Db.Order("concat(brand,' ',model_name) ASC").Preload("Categories").Where("brand = ?",brandName)
	}, premium)
}

func (c CarRepositoryImpl) SelectCarsByType(category string, premium bool) ([]entities.Car,error) {
	return selectCarsWithQuery(func() *gorm.DB {
		return  c.Db.Order("concat(brand,' ',model_name) ASC").Preload("Categories").Joins("join cars_categories_ass on cars_categories_ass.car_model_name = model_name").Where("car_category_name = ?", category)
	}, premium)
}
