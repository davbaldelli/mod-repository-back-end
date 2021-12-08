package controllers

import (
	"context"
	"firebase.google.com/go/v4/messaging"
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

func (f FirebaseControllerImpl) Notify(payload *messaging.Message) error {
	if _, err := f.Client.Send(f.Context, payload); err != nil {
		return err
	}
	return nil
}
