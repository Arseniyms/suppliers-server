package auth

import (
	"arseniyms/suppliers/server/internal/controllers/jwt"
	"encoding/json"
	"net/http"
	"os"
)

const LOGIN_PATH = "/login/"

func Login(w http.ResponseWriter, r *http.Request) {
	var admin Admin
	err := json.NewDecoder(r.Body).Decode(&admin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// err = ValidatePassword(admin.Password)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusUnauthorized)
	// 	return
	// }

	if admin.Password != os.Getenv("TEMP_PASSWORD") {
		http.Error(w, "Wrong password", http.StatusUnauthorized)
		return
	}

	tokenString, err := jwt.CreateToken(admin.Password)

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
