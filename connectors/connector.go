package connectors

import (
	"arseniyms/suppliers/server/auth"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUri       = "MONGO_URI"
	databaseName   = "DATABASE_NAME"
	collectionName = "COLLECTION_NAME"
)

const (
	COMPANIES_PATH = "/companies/"
)

type Success struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var client *mongo.Client

func ConnectToDB() *mongo.Client {
	mongoURI := os.Getenv(mongoUri)
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}
	log.Println(mongoURI)
	clientOptions := options.Client().ApplyURI(mongoURI)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

func GetSuccess(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully connected\nNow you can make requests"))
}

func GetData(w http.ResponseWriter, r *http.Request) {
	isAuth := auth.GetIsAuthentificated(r)

	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)
	if idStr != "" {
		getDataById(w, r)
		return
	}

	collection := getCollection()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.Background())

	var companies []Company
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

func getDataById(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)
	filter := bson.M{"extId": idStr}
	collection := getCollection()

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

func CreateItem(w http.ResponseWriter, r *http.Request) {
	var c Company
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.ExtId = uuid.New().String()
	fmt.Println("Company is: ", c)

	collection := getCollection()
	_, err := collection.InsertOne(context.Background(), c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	success := Success{Code: http.StatusCreated, Msg: "Company successfully created"}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(success)

}

func PatchItem(w http.ResponseWriter, r *http.Request) {
	var updateData map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id, ok := updateData["extId"].(string)
	if !ok || id == "" {
		http.Error(w, "Missing or invalid ID", http.StatusBadRequest)
		return
	}

	collection := getCollection()
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

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, COMPANIES_PATH)

	filter := bson.M{"extId": idStr}
	collection := getCollection()

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

func getCollection() *mongo.Collection {
	databaseNameEnv := os.Getenv(databaseName)
	collectionNameEnv := os.Getenv(collectionName)

	if databaseNameEnv == "" || collectionNameEnv == "" {
		log.Fatal("databaseNameEnv or collectionNameEnv environment variable not set")
	}

	return client.Database(databaseNameEnv).Collection(collectionNameEnv)
}

func validateCompany(c *Company, isAuth bool) {
	if !isAuth {
		c.Phones = ""
	}
}
