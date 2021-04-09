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
	var dbCars []db.Car

	if result := c.Db.Preload("Categories").Find(&dbCars); result.Error != nil{
		//return result.Error
	}
	var dbBrands []db.CarBrand
	if result := c.Db.Find(&dbBrands); result.Error != nil {
		//return result.Error
	}
	brandsNation := make(map[string]string)
	for _, brand := range dbBrands {
		brandsNation[brand.Name] = brand.Nation
	}

	for _, dbCar := range dbCars {
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: dbCar.DownloadLink},
			Brand: entities.CarBrand{
				Name:   dbCar.Brand,
				Nation: entities.Nation{Name: brandsNation[dbCar.Brand]},
			},
			ModelName:  dbCar.ModelName,
			Categories: allCategoriesToEntity(dbCar.Categories),
		})
	}
	fmt.Println(cars)
	return cars
}

func allCategoriesToEntity(dbCategories []db.CarCategory) []entities.CarCategory{
	var cats []entities.CarCategory
	for _,dbCat := range  dbCategories {
		cats = append(cats, entities.CarCategory{Name: dbCat.Name})
	}
	return cats
}


func (c CarRepositoryImpl) GetCarsByNation(nation string) []entities.Car {
	var cars []entities.Car
	var dbCars []db.Car

	if result := c.Db.Preload("Categories").Joins("join car_brands on cars.brand = car_brands.name").Where("car_brands.nation = ?",nation).Find(&dbCars); result.Error != nil{
		//return result.Error
	}
	var dbBrands []db.CarBrand
	if result := c.Db.Find(&dbBrands,"nation = ?",nation); result.Error != nil {
		//return result.Error
	}
	brandsNation := make(map[string]string)
	for _, brand := range dbBrands {
		brandsNation[brand.Name] = brand.Nation
	}

	for _, dbCar := range dbCars {
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: dbCar.DownloadLink},
			Brand: entities.CarBrand{
				Name:   dbCar.Brand,
				Nation: entities.Nation{Name: brandsNation[dbCar.Brand]},
			},
			ModelName:  dbCar.ModelName,
			Categories: allCategoriesToEntity(dbCar.Categories),
		})
	}
	fmt.Println(cars)
	return cars
}

func (c CarRepositoryImpl) GetCarByModel(model string) []entities.Car {
	var cars []entities.Car
	var dbCars []db.Car

	if result := c.Db.Preload("Categories").Find(&dbCars,"model_name = ?", model); result.Error != nil{
		//return result.Error
	}

	var brandsNames []string
	for _, car := range dbCars {
		brandsNames = append(brandsNames, car.Brand)
	}
	
	var dbBrands []db.CarBrand
	if result := c.Db.Find(&dbBrands, "name IN ?", brandsNames); result.Error != nil {
		//return result.Error
	}
	brandsNation := make(map[string]string)
	for _, brand := range dbBrands {
		brandsNation[brand.Name] = brand.Nation
	}

	for _, dbCar := range dbCars {
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: dbCar.DownloadLink},
			Brand: entities.CarBrand{
				Name:   dbCar.Brand,
				Nation: entities.Nation{Name: brandsNation[dbCar.Brand]},
			},
			ModelName:  dbCar.ModelName,
			Categories: allCategoriesToEntity(dbCar.Categories),
		})
	}
	fmt.Println(cars)
	return cars
}

func (c CarRepositoryImpl) GetCarsByBrand(brandName string) []entities.Car {
	var cars []entities.Car
	var dbCars []db.Car

	if result := c.Db.Preload("Categories").Find(&dbCars,"brand = ?",brandName); result.Error != nil{
		//return result.Error
	}
	dbBrand := db.CarBrand{Name: brandName}
	if result := c.Db.Find(&dbBrand); result.Error != nil {
		//return result.Error
	}

	for _, dbCar := range dbCars {
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: dbCar.DownloadLink},
			Brand: entities.CarBrand{
				Name:   dbCar.Brand,
				Nation: entities.Nation{Name: dbBrand.Nation},
			},
			ModelName:  dbCar.ModelName,
			Categories: allCategoriesToEntity(dbCar.Categories),
		})
	}
	fmt.Println(cars)
	return cars
}

func (c CarRepositoryImpl) GetCarsByType(category string) []entities.Car {
	
	var cars []entities.Car
	var dbCars []db.Car

	if result := c.Db.Preload("Categories").Joins("join cars_categories_ass on cars_categories_ass.car_model_name = model_name").Where("car_category_name = ?", category).Find(&dbCars); result.Error != nil{
		//return result.Error
	}
	var brandsNames []string
	for _, car := range dbCars {
		brandsNames = append(brandsNames, car.Brand)
	}

	var dbBrands []db.CarBrand
	if result := c.Db.Find(&dbBrands, "name IN ?", brandsNames); result.Error != nil {
		//return result.Error
	}
	brandsNation := make(map[string]string)
	for _, brand := range dbBrands {
		brandsNation[brand.Name] = brand.Nation
	}
	for _, dbCar := range dbCars {
		cars = append(cars, entities.Car{
			Mod: entities.Mod{DownloadLink: dbCar.DownloadLink},
			Brand: entities.CarBrand{
				Name:   dbCar.Brand,
				Nation: entities.Nation{Name: brandsNation[dbCar.Brand]},
			},
			ModelName:  dbCar.ModelName,
			Categories: allCategoriesToEntity(dbCar.Categories),
		})
	}
	fmt.Println(cars)
	return cars
}
