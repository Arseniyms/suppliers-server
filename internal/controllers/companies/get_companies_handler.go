package companies

import (
	"arseniyms/suppliers/server/internal/controllers/auth"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func GetCompanies(w http.ResponseWriter, r *http.Request) {
	isAuth := auth.GetIsAuthentificated(r)

	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)
	if idStr != "" {
		getCompanyById(w, r)
		return
	}

	collection := getCompaniesCollection()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var companies []Company = []Company{}
	for cursor.Next(context.Background()) {
		var c Company
		if err := cursor.Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		validateCompany(&c, isAuth)

		companies = append(companies, c)
	}

	if err = cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(companies)
}
