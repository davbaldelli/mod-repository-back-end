package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"net/http"
)

type UserHandlerImpl struct {
	UserCtrl controllers.LoginController
}

func (u UserHandlerImpl) POSTLogin(writer http.ResponseWriter, request *http.Request) {

	var reqUser entities.User

	decoder := json.NewDecoder(request.Body)

	if err := decoder.Decode(&reqUser); err != nil {
		respondError(writer, http.StatusBadRequest, fmt.Errorf("error converting post form to entiy: %v ", err))
		return
	}
	if user, err := u.UserCtrl.Login(reqUser.Username,reqUser.Password); err != nil{
		respondError(writer, http.StatusNotFound, err)
	} else {
		respondJSON(writer, http.StatusAccepted, user)
	}
}


