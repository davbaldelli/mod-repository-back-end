package controllers

import (
	"errors"
	"github.com/davide/ModRepository/models"
	"github.com/davide/ModRepository/repositories"
)

type TrackControllerImpl struct {
	repo repositories.TrackRepository
}

type trackParamCondition func(models.Track) bool

func (t TrackControllerImpl) getAllTracks() []models.Track {
	return t.repo.GetAllTracks()
}

func (t TrackControllerImpl) getTracksByNation(nationName string) []models.Track {
	return t.tracksResearchByParam(func(track models.Track) bool { return track.Location.Nation.Name == nationName })
}

func (t TrackControllerImpl) getTracksByLayoutType(category string) []models.Track {
	var tracks []models.Track
	for _, track := range t.repo.GetAllTracks() {
		for _, layout := range track.Layouts {
			if layout.LayoutType == category {
				tracks = append(tracks, track)
				break
			}
		}
	}
	return tracks
}

func (t TrackControllerImpl) getTrackByName(name string) (models.Track, error) {
	for _, track := range t.repo.GetAllTracks() {
		if track.Name == name {
			return track, nil
		}
	}
	return models.Track{}, errors.New("track " + name + "no found")
}

func (t TrackControllerImpl) addNewTrack(name string, downloadUrl string, layouts []models.Layout, locationName string, nation models.Nation) error {
	err := t.repo.AddNewTrack(models.Track{
		Mod:     models.Mod{DownloadLink: downloadUrl},
		Layouts: layouts,
		Location: models.Location{
			LocationName: locationName,
			Nation:       nation,
		},
		Name: name,
	})
	return err
}

func (t TrackControllerImpl) tracksResearchByParam(tpc trackParamCondition) []models.Track {
	var tracks []models.Track
	for _, track := range t.repo.GetAllTracks() {
		if tpc(track) {
			tracks = append(tracks, track)
		}
	}
	return tracks
}
