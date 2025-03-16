package password

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreatePassword(w http.ResponseWriter, r *http.Request) {
	var pass Admin
	if err := json.NewDecoder(r.Body).Decode(&pass); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("pass is: ", pass)

	collection := getPasswordCollection()
	_, err := collection.InsertOne(context.Background(), pass)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := Success{Code: http.StatusCreated, Msg: "Password successfully created"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success)
}
