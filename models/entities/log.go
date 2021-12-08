package entities

import "time"

type Action string

const (
	Update Action = "Update"
	Insert        = "Insert"
)

type CarLog struct {
	Id         uint      `json:"id"`
	Brand      string    `json:"brand"`
	Model      string    `json:"model"`
	Year       string    `json:"year"`
	Version    string    `json:"version"`
	Premium    bool      `json:"premium"`
	Action     Action    `json:"action"`
	HappenedAt time.Time `json:"happenedAt"`
}

type TrackLog struct {
	Id         uint      `json:"id"`
	Name       string    `json:"name"`
	Version    string    `json:"version"`
	Premium    bool      `json:"premium"`
	Action     Action    `json:"action"`
	HappenedAt time.Time `json:"happenedAt"`
}
