package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/davide/ModRepository/controllers"
	"github.com/davide/ModRepository/models/entities"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"
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

func GenerateJWT(username, role string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email" : username,
		"role" : role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("eskere"))

	if err != nil {
		return "", fmt.Errorf("Something Went Wrong: %s ", err.Error())
	}
	return tokenString, nil
}

func (u UserHandlerImpl)SignIn(w http.ResponseWriter, r *http.Request) {

	var authdetails entities.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error in request body: %v ", err))
		return
	}

	authuser, err := u.UserCtrl.Login(authdetails.Username, authdetails.Password)
	if err != nil {
		respondError(w, http.StatusNotFound, err)
		return
	}


	validToken, err := GenerateJWT(authuser.Username, string(authuser.Role))
	if err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error generating token: %v ", err))
		return
	}

	token := entities.Token{Username: authuser.Username, Role: string(authuser.Role), TokenString: validToken}

	respondJSON(w, http.StatusAccepted, token)
}