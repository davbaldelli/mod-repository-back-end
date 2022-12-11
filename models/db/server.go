package db

import "github.com/davide/ModRepository/models/entities"

type Server struct {
	Id          uint `gorm:"primaryKey"`
	Name        string
	Description string
	JoinLink    string
	Password    string
	Online      bool
	TrackId     uint
	Cars        []*Car `gorm:"many2many:server_cars;foreignKey:Id;joinForeignKey:ServerId;References:Id;joinReferences:CarId"`
}

// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer

func (s Server) ToEntity() entities.Server {
	var cars []uint
	for _, dbCar := range s.Cars {
		cars = append(cars, dbCar.Id)
	}
	return entities.Server{
		Id:          s.Id,
		Name:        s.Name,
		Description: s.Description,
		JoinLink:    s.JoinLink,
		Password:    s.Password,
		Online:      s.Online,
		Track:       s.TrackId,
		Cars:        cars,
	}
}
