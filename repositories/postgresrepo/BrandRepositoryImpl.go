package postgresrepo

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type BrandRepositoryImpl struct {
	Db *gorm.DB
}

func (b BrandRepositoryImpl) SelectAllBrands() ([]entities.CarBrand,error) {
	var dbBrands []db.CarBrand
	var brands []entities.CarBrand
	if result := b.Db.Find(&dbBrands); result.Error != nil {
		return nil,result.Error
	}
	for _, dbBrand := range dbBrands {
		brands = append(brands, entities.CarBrand{
			Name:   dbBrand.Name,
			Nation: entities.Nation{Name: dbBrand.Nation},
		})
	}
	return brands, nil
}

func (b BrandRepositoryImpl) SelectBrandsByNation(nation string) ([]entities.CarBrand, error) {
	var dbBrands []db.CarBrand
	var brands []entities.CarBrand
	if result := b.Db.Where("nation = ?", nation).Find(&dbBrands); result.Error != nil {
		return nil,result.Error
	}
	for _, dbBrand := range dbBrands {

		brands = append(brands, entities.CarBrand{
			Name:   dbBrand.Name,
			Nation: entities.Nation{Name: dbBrand.Nation},
		})
	}
	return brands, nil
}

func (b BrandRepositoryImpl) SelectBrandsByName(name string) ([]entities.CarBrand, error) {
	var dbBrands []db.CarBrand
	var brands []entities.CarBrand
	if result := b.Db.Where("name = ?", name).Find(&dbBrands); result.Error != nil {
		return nil,result.Error
	}
	for _, dbBrand := range dbBrands {
		brands = append(brands, entities.CarBrand{
			Name:   dbBrand.Name,
			Nation: entities.Nation{Name: dbBrand.Nation},
		})
	}
	return brands, nil
}

