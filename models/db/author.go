package db

type Author struct {
	Id     uint `gorm:"primarykey"`
	Name   string
	Link   string
	Cars   []Car   `gorm:"foreignKey:IdAuthor"`
	Tracks []Track `gorm:"foreignKey:IdAuthor"`
}
