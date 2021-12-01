package db

import (
	"time"
)

type ModModel struct {
	Id        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CarCategory struct {
	Id   uint `gorm:"primarykey"`
	Category string
	IdCar uint
}


func (CarCategory) TableName() string {
	return "car_categories"
}

type CarMods struct {
	ModModel
	Rating uint
	Version string
	DownloadLink string
	ModelName    string `gorm:"column:model"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"foreignKey:IdCar"`
	Transmission string
	Drivetrain   string
	Premium      bool
	Image        string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	Author       string
	AuthorLink   string
	Nation       string
	NationCode	string
}

type Car struct {
	ModModel
	Rating uint
	Version string
	DownloadLink string
	ModelName    string `gorm:"column:model"`
	Year        	int
	IdBrand      uint
	Categories   []CarCategory `gorm:"foreignKey:IdCar"`
	Transmission string
	Drivetrain   string
	Premium      bool
	Image        string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	IdAuthor     uint
}

type Author struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Link   string
	Cars   []Car   `gorm:"foreignKey:IdAuthor"`
	Tracks []Track `gorm:"foreignKey:IdAuthor"`
}

type Manufacturer struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	Cars     []Car `gorm:"foreignKey:IdBrand"`
	IdNation uint
}

type Nation struct {
	Id     uint           `gorm:"primarykey"`
	Name   string
	Code   string			`gorm:"type:varchar(6)"`
	Brands []Manufacturer `gorm:"foreignKey:IdNation"`
	Tracks []Track        `gorm:"foreignKey:IdNation"`
}

type Track struct {
	ModModel
	Version string
	DownloadLink string
	Name         string
	Layouts      []Layout `gorm:"foreignKey:IdTrack"`
	Location     string
	IdNation     uint
	Tags         []TrackTag `gorm:"foreignKey:IdTrack"`
	Year         uint
	Premium      bool
	Image        string
	IdAuthor     uint
	Rating uint
}

type TrackTag struct {
	Id      uint `gorm:"primarykey"`
	IdTrack uint
	Tag     string
}

type TrackMod struct {
	ModModel
	Version string
	DownloadLink string
	Name         string
	Layouts      []Layout   `gorm:"foreignKey:IdTrack"`
	Tags         []TrackTag `gorm:"foreignKey:IdTrack"`
	Location     string
	Nation       string
	NationCode string
	Year         uint
	Premium      bool
	Image        string
	Author       string
	AuthorLink   string
	Rating uint
}

type Layout struct {
	Id       uint `gorm:"primarykey"`
	Name     string
	LengthM  float32
	Category string
	IdTrack  uint
}

type User struct {
	Username string
	Password string
	Role     string
	Salt     string
}
