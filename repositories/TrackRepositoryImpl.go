package repositories

import (
	"context"
	"github.com/davide/ModRepository/models/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TrackRepositoryImpl struct {
	TrackCollection *mongo.Collection
}

func (t TrackRepositoryImpl) GetAllTracks() []entities.Track {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var tracks []entities.Track
	cursor, err := t.TrackCollection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &tracks); err != nil {
		panic(err)
	}
	return tracks
}

func (t TrackRepositoryImpl) GetTracksByNation(nationName string) []entities.Track {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var tracks []entities.Track
	cursor, err := t.TrackCollection.Find(ctx, bson.M{"location.nation.name": nationName})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &tracks); err != nil {
		panic(err)
	}
	return tracks
}

func (t TrackRepositoryImpl) GetTracksByLayoutType(layoutType string) []entities.Track {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var tracks []entities.Track
	cursor, err := t.TrackCollection.Find(ctx, bson.M{"layout.layouttype": layoutType})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &tracks); err != nil {
		panic(err)
	}
	return tracks
}

func (t TrackRepositoryImpl) GetTracksByName(name string) []entities.Track {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	var tracks []entities.Track
	cursor, err := t.TrackCollection.Find(ctx, bson.M{"name": name})
	if err != nil {
		panic(err)
	}
	if err = cursor.All(ctx, &tracks); err != nil {
		panic(err)
	}
	return tracks
}

func (t TrackRepositoryImpl) AddNewTrack(track entities.Track) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := t.TrackCollection.InsertOne(ctx, track)

	if err != nil {
		return err
	}
	return nil
}
