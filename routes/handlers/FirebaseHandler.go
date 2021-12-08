package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"net/http"
)

type FirebaseHandlerImpl struct {
	Ctrl controllers.FirebaseController
}

type SubscribeRequest struct {
	RegistrationToken string `json:"token"`
	Topic             string `json:"topic"`
}

func (f FirebaseHandlerImpl) SubscribeToTopic(writer http.ResponseWriter, request *http.Request) {
	regReq := SubscribeRequest{}

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&regReq); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}

	if err := f.Ctrl.RegisterToTopic(regReq.RegistrationToken, regReq.Topic); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error registrating token to topic: %v ", err))
		return
	}

	respondJSON(writer, http.StatusOK, "registration successful")

}
