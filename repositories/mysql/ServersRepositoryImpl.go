package mysql

import (
	models2 "github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories/models"
	"gorm.io/gorm"
)

type ServersRepositoryImpl struct {
	Db *gorm.DB
}

type serverCarsAssoc struct {
	CarId    uint
	ServerId uint
}

func (s ServersRepositoryImpl) UpdateServer(server models2.Server) error {

	dbServer := models.ServerFromEntity(server)

	if result := s.Db.Model(models.Server{}).Where("id = ?", dbServer.Id).Select("*").Omit("OutsideCars").Updates(&dbServer); result.Error != nil {
		return result.Error
	}

	var serverCars []serverCarsAssoc

	for _, carId := range server.Cars {
		serverCars = append(serverCars, serverCarsAssoc{
			CarId:    carId,
			ServerId: server.Id,
		})
	}

	if result := s.Db.Model(&models.Server{Id: dbServer.Id}).Association("Cars").Clear(); result != nil {
		return result
	}

	if len(serverCars) > 0 {
		if result := s.Db.Table("server_cars").Create(&serverCars); result.Error != nil {
			return result.Error
		}
	}

	var outsideCars []models.OutsideMod

	for _, outsideCar := range server.OutsideCars {
		outsideCars = append(outsideCars, models.OutsideModFromEntity(outsideCar, dbServer.Id))
	}

	if result := s.Db.Where("server_id = ?", dbServer.Id).Delete(&models.OutsideMod{}); result.Error != nil {
		return result.Error
	}

	if len(outsideCars) > 0 {
		if result := s.Db.Debug().Model(&models.OutsideMod{}).Omit("Id").Create(&outsideCars); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s ServersRepositoryImpl) AddServer(server models2.Server) error {

	dbServer := models.ServerFromEntity(server)

	if result := s.Db.Model(models.Server{}).Omit("Cars", "OutsideCars").Create(&dbServer); result.Error != nil {
		return result.Error
	}

	var serverCars []serverCarsAssoc
	serverCars = make([]serverCarsAssoc, 0)

	for _, carId := range server.Cars {
		serverCars = append(serverCars, serverCarsAssoc{
			CarId:    carId,
			ServerId: dbServer.Id,
		})
	}

	if len(serverCars) > 0 {
		if result := s.Db.Table("server_cars").Create(&serverCars); result.Error != nil {
			return result.Error
		}
	}

	var outsideCars []models.OutsideMod
	outsideCars = make([]models.OutsideMod, 0)

	for _, outsideCar := range server.OutsideCars {
		outsideCars = append(outsideCars, models.OutsideModFromEntity(outsideCar, dbServer.Id))
	}

	if len(outsideCars) > 0 {
		if result := s.Db.Model(&models.OutsideMod{}).Omit("Id").Create(&outsideCars); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (s ServersRepositoryImpl) DeleteServer(server models2.Server) error {
	if result := s.Db.Where("id = ?", server.Id).Delete(&models.Server{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func (s ServersRepositoryImpl) GetAllServers() ([]models2.Server, error) {
	var servers []models2.Server
	var dbServers []models.Server
	if result := s.Db.Model(models.Server{}).Preload("Cars").Preload("OutsideCars").Find(&dbServers); result.Error != nil {
		return nil, result.Error
	}
	for _, server := range dbServers {
		servers = append(servers, server.ToEntity())
	}
	return servers, nil
}
