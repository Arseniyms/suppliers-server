package password

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

func IsPasswordEmpty(w http.ResponseWriter, r *http.Request) {
	collection := getPasswordCollection()

	count, err := collection.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if count == 0 {
		http.Error(w, "no password", http.StatusNoContent)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
