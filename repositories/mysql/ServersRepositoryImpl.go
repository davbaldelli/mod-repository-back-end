package mysql

import (
	"github.com/davide/ModRepository/models/db"
	"github.com/davide/ModRepository/models/entities"
	"gorm.io/gorm"
)

type ServersRepositoryImpl struct {
	Db *gorm.DB
}

type serverCarsAssoc struct {
	CarId    uint
	ServerId uint
}

func (s ServersRepositoryImpl) UpdateServer(server entities.Server) error {

	if result := s.Db.Model(db.Server{}).Where("id = ?", server.Id).Updates(&server); result.Error != nil {
		return result.Error
	}

	var serverCars []serverCarsAssoc

	for _, carId := range server.Cars {
		serverCars = append(serverCars, serverCarsAssoc{
			CarId:    carId,
			ServerId: server.Id,
		})
	}

	if result := s.Db.Model(&db.Server{Id: server.Id}).Association("Cars").Clear(); result != nil {
		return result
	}

	if result := s.Db.Table("server_cars").Create(&serverCars); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s ServersRepositoryImpl) AddServer(server entities.Server) error {

	if result := s.Db.Model(db.Server{}).Omit("Cars").Create(&server); result.Error != nil {
		return result.Error
	}

	var serverCars []serverCarsAssoc

	for _, carId := range server.Cars {
		serverCars = append(serverCars, serverCarsAssoc{
			CarId:    carId,
			ServerId: server.Id,
		})
	}

	if result := s.Db.Table("server_cars").Create(&serverCars); result.Error != nil {
		return result.Error
	}

	return nil
}

func (s ServersRepositoryImpl) GetAllServers() ([]entities.Server, error) {
	var servers []entities.Server
	var dbServers []db.Server
	if result := s.Db.Model(db.Server{}).Preload("Cars").Find(&dbServers); result.Error != nil {
		return nil, result.Error
	}
	for _, server := range dbServers {
		servers = append(servers, server.ToEntity())
	}
	return servers, nil
}
