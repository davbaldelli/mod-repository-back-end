package models

import "time"

type ModModel struct {
	Id           uint `gorm:"primaryKey"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Rating       uint
	Version      string
	DownloadLink string
	Source       string
	Premium      bool
	Personal     bool
	IdAuthor     uint
	Official     bool
}
