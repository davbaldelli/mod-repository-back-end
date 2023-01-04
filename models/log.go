package models

import "time"

type Action string

const (
	Update Action = "Update"
	Insert        = "Insert"
)

type CarLog struct {
	LogId      uint      `json:"logId"`
	CarId      uint      `json:"carId"`
	Brand      string    `json:"brand"`
	Model      string    `json:"model"`
	Year       string    `json:"year"`
	Version    string    `json:"version"`
	Premium    bool      `json:"premium"`
	Action     Action    `json:"action"`
	HappenedAt time.Time `json:"happenedAt"`
}

type TrackLog struct {
	LogId      uint      `json:"logId"`
	TrackId    uint      `json:"trackId"`
	Name       string    `json:"name"`
	Year       string    `json:"year"`
	Version    string    `json:"version"`
	Premium    bool      `json:"premium"`
	Action     Action    `json:"action"`
	HappenedAt time.Time `json:"happenedAt"`
}
