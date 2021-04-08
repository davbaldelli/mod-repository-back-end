package postgresrepo

import (
	"fmt"
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CarRepositoryImpl struct {
	Db *gorm.DB
}

type rawCar struct {
	Brand        string
	DownloadLink string
	ModelName    string
	Nation       string
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
	var cars []entities.Car
	var result []rawCar
	c.Db.Model(&db.Car{}).Select("*").Joins("join car_brands on cars.brand = car_brands.name ").Scan(&result)
	for _, row := range result {
		var dbCategories []db.CarCategory
		var categories []entities.CarCategory
		c.Db.Model(&db.CarCategory{}).Select("car_categories.*").Joins("join cars_categories_ass on cars_categories_ass.car_category_name = name").Where("car_model_name = ?", row.ModelName).Scan(&dbCategories)
		for _, dbCategory := range dbCategories {
			categories = append(categories, entities.CarCategory{Name: dbCategory.Name})
		}
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: row.DownloadLink},
			Brand: entities.CarBrand{
				Name:   row.Brand,
				Nation: entities.Nation{Name: row.Nation},
			},
			ModelName:  row.ModelName,
			Categories: categories,
		})
	}
	fmt.Println(result)
	return cars
}

func (c CarRepositoryImpl) GetCarsByNation(nation string) []entities.Car {
	var cars []entities.Car
	var result []rawCar
	c.Db.Model(&db.Car{}).Select("*").Joins("join car_brands on cars.brand = car_brands.name ").Where("nation = ?", nation).Scan(&result)
	for _, row := range result {
		var dbCategories []db.CarCategory
		var categories []entities.CarCategory
		c.Db.Model(&db.CarCategory{}).Select("car_categories.*").Joins("join cars_categories_ass on cars_categories_ass.car_category_name = name").Where("car_model_name = ?", row.ModelName).Scan(&dbCategories)
		for _, dbCategory := range dbCategories {
			categories = append(categories, entities.CarCategory{Name: dbCategory.Name})
		}
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: row.DownloadLink},
			Brand: entities.CarBrand{
				Name:   row.Brand,
				Nation: entities.Nation{Name: row.Nation},
			},
			ModelName:  row.ModelName,
			Categories: categories,
		})
	}
	fmt.Println(result)
	return cars
}

func (c CarRepositoryImpl) GetCarByModel(model string) []entities.Car {
	var cars []entities.Car
	var result []rawCar
	c.Db.Model(&db.Car{}).Select("*").Joins("join car_brands on cars.brand = car_brands.name ").Where("model_name = ?", model).Scan(&result)
	for _, row := range result {
		var dbCategories []db.CarCategory
		var categories []entities.CarCategory
		c.Db.Model(&db.CarCategory{}).Select("car_categories.*").Joins("join cars_categories_ass on cars_categories_ass.car_category_name = name").Where("car_model_name = ?", row.ModelName).Scan(&dbCategories)
		for _, dbCategory := range dbCategories {
			categories = append(categories, entities.CarCategory{Name: dbCategory.Name})
		}
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: row.DownloadLink},
			Brand: entities.CarBrand{
				Name:   row.Brand,
				Nation: entities.Nation{Name: row.Nation},
			},
			ModelName:  row.ModelName,
			Categories: categories,
		})
	}
	fmt.Println(result)
	return cars
}

func (c CarRepositoryImpl) GetCarsByBrand(brand string) []entities.Car {
	var cars []entities.Car
	var result []rawCar
	c.Db.Model(&db.Car{}).Select("*").Joins("join car_brands on cars.brand = car_brands.name ").Where("brand = ?", brand).Scan(&result)
	for _, row := range result {
		var dbCategories []db.CarCategory
		var categories []entities.CarCategory
		c.Db.Model(&db.CarCategory{}).Select("car_categories.*").Joins("join cars_categories_ass on cars_categories_ass.car_category_name = name").Where("car_model_name = ?", row.ModelName).Scan(&dbCategories)
		for _, dbCategory := range dbCategories {
			categories = append(categories, entities.CarCategory{Name: dbCategory.Name})
		}
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: row.DownloadLink},
			Brand: entities.CarBrand{
				Name:   row.Brand,
				Nation: entities.Nation{Name: row.Nation},
			},
			ModelName:  row.ModelName,
			Categories: categories,
		})
	}
	fmt.Println(result)
	return cars
}

func (c CarRepositoryImpl) GetCarsByType(category string) []entities.Car {
	var cars []entities.Car
	var result []rawCar
	c.Db.Model(&db.Car{}).Select("cars.* ", "car_brands.*").Joins("join car_brands on cars.brand = car_brands.name ").Joins("join cars_categories_ass on cars_categories_ass.car_model_name = model_name").Where("car_category_name = ?", category).Scan(&result)
	for _, row := range result {
		var dbCategories []db.CarCategory
		var categories []entities.CarCategory
		c.Db.Model(&db.CarCategory{}).Select("car_categories.*").Joins("join cars_categories_ass on cars_categories_ass.car_category_name = name").Where("car_model_name = ?", row.ModelName).Scan(&dbCategories)
		for _, dbCategory := range dbCategories {
			categories = append(categories, entities.CarCategory{Name: dbCategory.Name})
		}
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: row.DownloadLink},
			Brand: entities.CarBrand{
				Name:   row.Brand,
				Nation: entities.Nation{Name: row.Nation},
			},
			ModelName:  row.ModelName,
			Categories: categories,
		})
	}
	fmt.Println(result)
	return cars
}
