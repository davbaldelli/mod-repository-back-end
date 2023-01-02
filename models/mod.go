package models

import "time"

type Mod struct {
	Id           uint      `json:"id"`
	DownloadLink string    `json:"downloadLink"`
	Source       string    `json:"source"`
	Premium      bool      `json:"premium"`
	Personal     bool      `json:"personal"`
	Images       []Image   `json:"images"`
	Author       Author    `json:"author"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Rating       uint      `json:"rating"`
	Version      string    `json:"version"`
	Official     bool      `json:"official"`
}

type Author struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
