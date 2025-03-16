package auth

import (
	"arseniyms/suppliers/server/internal/controllers/jwt"
	"errors"
	"net/http"
	"strings"
)

func ValidateConnection(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	if len(tokenString) == 0 {
		return "", errors.New("no token")
	}

	token, err := jwt.ValidateJWT(tokenString)
	if err != nil {
		return "", errors.New("unauthorized")
	}
	return token.Raw, nil
}
