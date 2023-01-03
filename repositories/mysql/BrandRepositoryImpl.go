package mysql

import (
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/entities"
	"gorm.io/gorm"
)

type BrandRepositoryImpl struct {
	Db *gorm.DB
}

func (b BrandRepositoryImpl) SelectAllBrands() ([]models2.CarBrand, error) {
	return b.selectBrandsWithQuery(func(brands *[]entities.Manufacturer) *gorm.DB {
		return b.Db.Order("name ASC").Find(&brands)
	})
}

func (b BrandRepositoryImpl) selectBrandsWithQuery(query selectFromBrandsQuery) ([]models2.CarBrand, error) {
	var dbBrands []entities.Manufacturer
	var brands []models2.CarBrand
	if result := query(&dbBrands); result.Error != nil {
		return nil, result.Error
	}
	for _, dbBrand := range dbBrands {
		nation := entities.Nation{Id: dbBrand.IdNation}
		if res2 := b.Db.Find(&nation); res2.Error != nil {
			return nil, res2.Error
		} else {
			brands = append(brands, dbBrand.ToEntity(nation))
		}
	}
	return brands, nil
}
