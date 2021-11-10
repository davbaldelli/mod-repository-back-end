package handlers

import (
	"fmt"
	"github.com/davide/ModRepository/models/entities"
	"github.com/golang-jwt/jwt"
	"net/http"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] == nil {
			respondError(w, http.StatusUnauthorized, fmt.Errorf("token not found"))
			return
		}
		var mySigningKey = []byte("eskere")

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error in parsing")
			}
			return mySigningKey, nil
		})

		if err != nil {
			respondError(w, http.StatusUnauthorized, fmt.Errorf("your Token has been expired: %v", err.Error()))
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			switch claims["role"] {
			case "admin":
				{
					r.Header.Set("Role", string(entities.Admin))
					next.ServeHTTP(w, r)
					return
				}

			case "premium":
				{
					r.Header.Set("Role", string(entities.Premium))
					next.ServeHTTP(w, r)
					return
				}
			case "base":
				{
					r.Header.Set("Role", string(entities.Base))
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		respondError(w, http.StatusUnauthorized, fmt.Errorf("you have no authorization"))
	}
}

func IsAllowed(next http.HandlerFunc, allowedRoles []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if contains(allowedRoles, r.Header["Role"][0]) {
			next.ServeHTTP(w, r)
		} else {
			respondError(w, http.StatusForbidden, fmt.Errorf("you are not allowed to use this resource"))
		}
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
