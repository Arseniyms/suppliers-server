package companies

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func CreateCompanyArray(w http.ResponseWriter, r *http.Request) {
	var cArr []Company
	if err := json.NewDecoder(r.Body).Decode(&cArr); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range cArr {
		cArr[i].ExtId = uuid.New().String()
	}

	var documents []interface{}
	for _, company := range cArr {
		documents = append(documents, company)
	}

	collection := getCompaniesCollection()
	_, err := collection.InsertMany(context.Background(), documents)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := Success{Code: http.StatusCreated, Msg: "Companies successfully created"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success)
}
