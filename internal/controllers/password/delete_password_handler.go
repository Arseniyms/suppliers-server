package password

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func DeletePassword(w http.ResponseWriter, r *http.Request) {
	collection := getPasswordCollection()

	result, err := collection.DeleteMany(context.Background(), bson.D{})

	if err != nil {
		http.Error(w, "Error deleting item", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "No document found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	success := Success{Code: http.StatusOK, Msg: "Passwords successfully deleted"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(success)
}
