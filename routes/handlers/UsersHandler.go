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

func GenerateJWT(username, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"email":      username,
		"role":       role,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("eskere"))

	if err != nil {
		return "", fmt.Errorf("Something Went Wrong: %s ", err.Error())
	}
	return tokenString, nil
}

func (u UserHandlerImpl) LogIn(w http.ResponseWriter, r *http.Request) {

	var authdetails entities.Authentication
	err := json.NewDecoder(r.Body).Decode(&authdetails)
	if err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error in request body: %v ", err))
		return
	}

	authuser, err := u.UserCtrl.Login(authdetails.Username, authdetails.Password)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err)
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

func (u UserHandlerImpl) SignIn(w http.ResponseWriter, r *http.Request) {

	var user entities.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		respondError(w, http.StatusBadRequest, fmt.Errorf("error in request body: %v ", err))
		return
	}

	newUser, err := u.UserCtrl.SignIn(user.Username, user.Password, user.Role)
	if err != nil {
		respondError(w, http.StatusUnauthorized, err)
		return
	}

	validToken, err := GenerateJWT(newUser.Username, string(newUser.Role))
	if err != nil {
		respondError(w, http.StatusInternalServerError, fmt.Errorf("error generating token: %v ", err))
		return
	}

	token := entities.Token{Username: newUser.Username, Role: string(newUser.Role), TokenString: validToken}

	respondJSON(w, http.StatusAccepted, token)
}
