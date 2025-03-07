package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

const LOGIN_PATH = "/login/"

type LoginSuccess struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Token string `json:"token"`
}

type Admin struct {
	Password string `json:password`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if admin.Password != "1234" {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	tokenString, err := CreateToken(admin.Password)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := LoginSuccess{
		Code:  http.StatusOK,
		Msg:   "Admin successfully logged in",
		Token: tokenString,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	if len(tokenString) == 0 {
		http.Error(w, "No token", http.StatusUnauthorized)
		return
	}

	token, err := ValidateJWT(tokenString)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	success := LoginSuccess{Code: http.StatusOK, Msg: "Welcome ", Token: token.Raw}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}

func ValidateConnection(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	if len(tokenString) == 0 {
		return "", errors.New("no token")
	}

	token, err := ValidateJWT(tokenString)
	if err != nil {
		return "", errors.New("unauthorized")
	}
	return token.Raw, nil
}
