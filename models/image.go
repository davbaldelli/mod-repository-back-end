package models

type Image struct {
	Id       uint   `json:"id"`
	Url      string `json:"url"`
	Favorite bool   `json:"favorite"`
}
