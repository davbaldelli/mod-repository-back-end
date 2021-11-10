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
	Name string
	Cars []Car `gorm:"many2many:cars_categories_ass;joinForeignKey:id_category"`
}
type CarMods struct {
	ModModel
	DownloadLink string
	ModelName    string `gorm:"column:model"`
	Year         uint
	Brand        string
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;joinForeignKey:id_car"`
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
	ModModel
	DownloadLink string
	ModelName    string `gorm:"column:model"`
	Year         uint
	IdBrand      uint
	Categories   []CarCategory `gorm:"many2many:cars_categories_ass;joinForeignKey:id_car"`
	Transmission string
	Drivetrain   string
	Premium      bool
	Image        string
	BHP          uint
	Torque       uint
	Weight       uint
	TopSpeed     uint
	IdAuthor 	uint
}

type Author struct {
	Id   uint `gorm:"primarykey"`
	Name string
	Link string
	Cars []Car    `gorm:"foreignKey:IdAuthor"`
	Tracks []Track `gorm:"foreignKey:IdAuthor"`
}

type Manufacturer struct {
	Id        uint `gorm:"primarykey"`
	Name   string
	Cars   []Car `gorm:"foreignKey:IdBrand"`
	IdNation uint
}

type Nation struct {
	Id        uint `gorm:"primarykey"`
	Name   string     `gorm:"primaryKey"`
	Brands []Manufacturer `gorm:"foreignKey:IdNation"`
	Tracks []Track `gorm:"foreignKey:IdNation"`
}

type Track struct {
	ModModel
	DownloadLink string
	Name         string
	Layouts      []Layout `gorm:"foreignKey:IdTrack"`
	Location     string
	IdNation       uint
	//Tags 		 pq.StringArray `gorm:"type:track_tag[]"`
	Year 		 uint
	Premium 	 bool
	Image string
	IdAuthor uint
}

type TrackMod struct {
	ModModel
	DownloadLink string
	Name         string
	Layouts      []Layout `gorm:"foreignKey:IdTrack"`
	Location     string
	Nation       string
	Year 		 uint
	Premium 	 bool
	Image string
	Author string
	AuthorLink string
}

type Layout struct {
	Id        uint `gorm:"primarykey"`
	Name     string
	LengthM  float32
	Category string
	IdTrack  uint
}

type User struct {
	Username string
	Password string
	Role string
	Salt string
}

























