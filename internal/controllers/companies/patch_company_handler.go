package companies

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func PatchCompany(w http.ResponseWriter, r *http.Request) {
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, ok := updateData["extId"].(string)
	if !ok || id == "" {
		http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
		return
	}

	collection := getCompaniesCollection()
	delete(updateData, "extId")
	filter := bson.M{"extId": id}
	update := bson.M{"$set": updateData}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if result.MatchedCount == 0 {
		http.Error(w, "No item found with ID", http.StatusNotFound)
		return
	}

	success := Success{Code: http.StatusOK, Msg: "Company successfully updated"}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}
