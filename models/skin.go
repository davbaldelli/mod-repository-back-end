package models

type Skin struct {
	Id           uint   `json:"id"`
	Name         string `json:"name"`
	DownloadLink string `json:"downloadLink"`
	ImageUrl     string `json:"imageUrl"`
	CarId        uint   `json:"carId"`
}
