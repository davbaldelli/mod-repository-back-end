package controllers

import (
	"context"
	"firebase.google.com/go/v4/messaging"
	"fmt"
	"github.com/davide/ModRepository/models"
)

type FirebaseControllerImpl struct {
	Client  *messaging.Client
	Context context.Context
}

func (f FirebaseControllerImpl) RegisterToTopic(token string, topic string) error {
	if _, err := f.Client.SubscribeToTopic(f.Context, []string{token}, topic); err != nil {
		return err
	}
	return nil
}

func (f FirebaseControllerImpl) NotifyCarUpdated(car models.Car) error {
	payload := &messaging.Message{
		Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
			{Action: "car_updated", Title: "Check It Out!"},
		}}},
		Notification: &messaging.Notification{Title: fmt.Sprintf("%v %v has been updated", car.Brand.Name, car.ModelName), Body: "A resource has been updated", ImageURL: "https://imgur.com/0GuN24g"},
		Topic:        "modsUpdates",
	}
	if _, err := f.Client.Send(f.Context, payload); err != nil {
		return err
	}
	return nil
}

func (f FirebaseControllerImpl) NotifyCarAdded(car models.Car) error {
	payload := &messaging.Message{
		Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
			{Action: "car_added", Title: "Check It Out!"},
		}}},
		Notification: &messaging.Notification{Title: fmt.Sprintf("%v %v has been added to repository", car.Brand.Name, car.ModelName), Body: "A resource has been added", ImageURL: "https://imgur.com/0GuN24g"},
		Topic:        "modsUpdates",
	}
	if _, err := f.Client.Send(f.Context, payload); err != nil {
		return err
	}
	return nil
}

func (f FirebaseControllerImpl) NotifyTrackAdded(track models.Track) error {
	payload := &messaging.Message{
		Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
			{Action: "track_added", Title: "Check It Out!"},
		}}},
		Notification: &messaging.Notification{Title: fmt.Sprintf("%v has been added to repository", track.Name), Body: "A resource has been added", ImageURL: "https://imgur.com/0GuN24g"},
		Topic:        "modsUpdates",
	}
	if _, err := f.Client.Send(f.Context, payload); err != nil {
		return err
	}
	return nil
}
func (f FirebaseControllerImpl) NotifyTrackUpdated(track models.Track) error {
	payload := &messaging.Message{
		Webpush: &messaging.WebpushConfig{Notification: &messaging.WebpushNotification{Actions: []*messaging.WebpushNotificationAction{
			{Action: "track_updated", Title: "Check It Out!"},
		}}},
		Notification: &messaging.Notification{Title: fmt.Sprintf("%v has been updated", track.Name), Body: "A resource has been updated", ImageURL: "https://imgur.com/0GuN24g"},
		Topic:        "modsUpdates",
	}
	if _, err := f.Client.Send(f.Context, payload); err != nil {
		return err
	}
	return nil
}
