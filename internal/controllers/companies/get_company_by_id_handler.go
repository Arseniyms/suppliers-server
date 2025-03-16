package companies

import (
	"arseniyms/suppliers/server/internal/controllers/auth"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func getCompanyById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)
	filter := bson.M{"extId": idStr}
	collection := getCompaniesCollection()

	var result Company
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err == mongo.ErrNoDocuments {
		http.Error(w, "No found items", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error finding item", http.StatusInternalServerError)
		return
	}

	validateCompany(&result, auth.GetIsAuthentificated(r))

	json.NewEncoder(w).Encode(result)
}

func GetCompanyById(w http.ResponseWriter, idStr string) (*Company, error) {
	filter := bson.M{"extId": idStr}
	collection := getCompaniesCollection()

	var result Company
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
