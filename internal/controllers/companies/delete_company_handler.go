package companies

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
)

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)

	filter := bson.M{"extId": idStr}
	collection := getCompaniesCollection()

	result, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "No document found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	success := Success{Code: http.StatusOK, Msg: "Company successfully deleted"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}
