package mysql

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type BrandRepositoryImpl struct {
	Db *gorm.DB
}

func (b BrandRepositoryImpl) SelectAllBrands(premium bool) ([]entities.CarBrand, error) {
	return b.selectBrandsWithQuery(func(brands *[]db.Manufacturer) *gorm.DB {
		return b.Db.Order("name ASC").Find(&brands)
	})
}

func (b BrandRepositoryImpl) selectBrandsWithQuery(query selectFromBrandsQuery) ([]entities.CarBrand, error) {
	var dbBrands []db.Manufacturer
	var brands []entities.CarBrand
	if result := query(&dbBrands); result.Error != nil {
		return nil, result.Error
	}
	for _, dbBrand := range dbBrands {
		nation := db.Nation{Id: dbBrand.IdNation}
		if res2 := b.Db.Find(&nation); res2.Error != nil {
			return nil, res2.Error
		} else {
			brands = append(brands, dbBrand.ToEntity(nation))
		}
	}
	return brands, nil
}
