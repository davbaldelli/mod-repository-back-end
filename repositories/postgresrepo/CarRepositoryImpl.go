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

type selectFromCarsQuery func(*[]db.Car) *gorm.DB
type selectFromBrandsQuery func(*[]db.CarBrand) *gorm.DB

func selectCarsWithQuery(carsQuery selectFromCarsQuery, brandsQuery selectFromBrandsQuery) ([]entities.Car, error){
	var cars []entities.Car
	var dbCars []db.Car

	if result := carsQuery(&dbCars); result.Error != nil{
		return nil,result.Error
	}
	var dbBrands []db.CarBrand
	if result := brandsQuery(&dbBrands); result.Error != nil {
		return nil,result.Error
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
			Drivetrain: entities.Drivetrain(dbCar.Drivetrain),
			GearType: entities.GearType(dbCar.GearType),
		})
	}
	return cars,nil
}

func (c CarRepositoryImpl) SelectAllCarCategories() ([]entities.CarCategory, error) {
	var categories []db.CarCategory
	if result := c.Db.Find(&categories) ; result.Error != nil{
		return  nil, result.Error
	}
	return allCategoriesToEntity(categories), nil
}

func (c CarRepositoryImpl) InsertCar(car entities.Car) error {
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

func (c CarRepositoryImpl) SelectAllCars() ([]entities.Car,error) {
	return selectCarsWithQuery(func(cars *[]db.Car) *gorm.DB {
		return c.Db.Order("model_name ASC").Preload("Categories").Find(&cars)
	}, func(brands *[]db.CarBrand) *gorm.DB {
		return c.Db.Find(&brands)
	})
}

func allCategoriesToEntity(dbCategories []db.CarCategory) []entities.CarCategory{
	var cats []entities.CarCategory
	for _,dbCat := range  dbCategories {
		cats = append(cats, entities.CarCategory{Name: dbCat.Name})
	}
	return cats
}


func (c CarRepositoryImpl) SelectCarsByNation(nation string) ([]entities.Car,error) {
	return selectCarsWithQuery(func(cars *[]db.Car) *gorm.DB {
		return c.Db.Order("model_name ASC").Preload("Categories").Joins("join car_brands on cars.brand = car_brands.name").Where("car_brands.nation = ?",nation).Find(&cars)
	}, func(brands *[]db.CarBrand) *gorm.DB {
		return c.Db.Find(&brands,"nation = ?",nation)
	})

}

func (c CarRepositoryImpl) SelectCarsByModelName(model string) ([]entities.Car,error) {
	return selectCarsWithQuery(func(cars *[]db.Car) *gorm.DB {
		return c.Db.Order("model_name ASC").Preload("Categories").Find(&cars,"model_name = ?", model).Find(&cars)
	}, func(brands *[]db.CarBrand) *gorm.DB {
		return c.Db.Joins("right join cars on cars.brand = car_brands.name").Find(&brands,"cars.model_name = ?",model)
	})
}

func (c CarRepositoryImpl) SelectCarsByBrand(brandName string) ([]entities.Car,error) {
	return selectCarsWithQuery(func(cars *[]db.Car) *gorm.DB {
		return c.Db.Order("model_name ASC").Preload("Categories").Find(&cars,"brand = ?",brandName)
	}, func(brands *[]db.CarBrand) *gorm.DB {
		return  c.Db.Find(&brands,"name = ?", brandName)
	})
}

func (c CarRepositoryImpl) SelectCarsByType(category string) ([]entities.Car,error) {
	
	var cars []entities.Car
	var dbCars []db.Car

	if result := c.Db.Order("model_name ASC").Preload("Categories").Joins("join cars_categories_ass on cars_categories_ass.car_model_name = model_name").Where("car_category_name = ?", category).Find(&dbCars); result.Error != nil{
		return nil,result.Error
	}
	var brandsNames []string
	for _, car := range dbCars {
		brandsNames = append(brandsNames, car.Brand)
	}

	var dbBrands []db.CarBrand
	if result := c.Db.Find(&dbBrands, "name IN ?", brandsNames); result.Error != nil {
		return nil,result.Error
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
	return cars, nil
}
