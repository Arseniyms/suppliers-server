package auth

import (
	"arseniyms/suppliers/server/internal/controllers/jwt"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func ProtectedEndpoint(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	if len(tokenString) == 0 {
		http.Error(w, "No token", http.StatusUnauthorized)
		return
	}

	token, err := jwt.ValidateJWT(tokenString)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	success := LoginSuccess{Code: http.StatusOK, Msg: "Welcome ", Token: token.Raw}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}
