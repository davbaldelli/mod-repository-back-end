package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc{
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
					r.Header.Set("Role", "admin")
					next.ServeHTTP(w, r)
					return
				}

			case "premium":
				{
					r.Header.Set("Role", "premium")
					next.ServeHTTP(w, r)
					return
				}
			case "base":
				{
					r.Header.Set("Role", "base")
					next.ServeHTTP(w, r)
					return
				}
			}
		}
		respondError(w, http.StatusUnauthorized, fmt.Errorf("you have no authorization"))
	}
}
