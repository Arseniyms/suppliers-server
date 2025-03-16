package password

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	mongoUri               = "MONGO_URI"
	databaseName           = "DATABASE_NAME"
	passwordCollectionName = "PASSWORD_COLLECTION_NAME"
	PASSWORD_PATH          = "/password/"
)

var client *mongo.Client

func SetClient(c *mongo.Client) {
	client = c
}

func getPasswordCollection() *mongo.Collection {
	databaseNameEnv := os.Getenv(databaseName)
	collectionNameEnv := os.Getenv(passwordCollectionName)

	if databaseNameEnv == "" || collectionNameEnv == "" {
		log.Fatal("databaseNameEnv or passwordCollectionName environment variable not set")
	}

	return client.Database(databaseNameEnv).Collection(collectionNameEnv)
}
