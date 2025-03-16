package companies

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var c Company
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.ExtId = uuid.New().String()
	fmt.Println("Company is: ", c)

	collection := getCompaniesCollection()
	_, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := Success{Code: http.StatusCreated, Msg: "Company successfully created"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success)
}
