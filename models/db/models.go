package db

import (
	"github.com/lib/pq"
)

type CarCategory struct {
	Name string `gorm:"primaryKey:not null"`
}
type CarMods struct {
	DownloadLink string
	ModelName    string `gorm:"primaryKey"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;joinForeignKey:car_model_name"`
	Transmission string
	Drivetrain   string
	Premium      bool
	Image        string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Author 		string
	AuthorLink string
	Nation string
}

type Car struct {
	DownloadLink string
	ModelName    string `gorm:"primaryKey"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;"`
	Transmission string
	Drivetrain   string
	Premium      bool
	Image        string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Author 		string
}

type Author struct {
	Name string `gorm:"primaryKey"`
	Link string
	Cars []CarMods `gorm:"foreignKey:Author"`
	Tracks []Track `gorm:"foreignKey:Author"`
}

type CarBrand struct {
	Name   string    `gorm:"primaryKey"`
	Cars   []CarMods `gorm:"foreignKey:Brand"`
	Nation string
}

type Nation struct {
	Name   string     `gorm:"primaryKey"`
	Brands []CarBrand `gorm:"foreignKey:Nation"`
	Tracks []Track    `gorm:"foreignKey:Nation"`
}

type Track struct {
	DownloadLink string
	Name         string   `gorm:"primaryKey"`
	Layouts      []Layout `gorm:"foreignKey:Track"`
	Location     string
	Nation       string
	Tags 		 pq.StringArray `gorm:"type:track_tag[]"`
	Year 		 uint
	Premium 	 bool
	Image string
	Author string
}

type Layout struct {
	Name     string `gorm:"primaryKey"`
	LengthM float32
	Category string
	Track    string `gorm:"primaryKey"`
}

type User struct {
	Username string `gorm:"primaryKey"`
	Password string
	IsAdmin bool
}

























