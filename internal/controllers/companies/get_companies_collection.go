package companies

import (
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

func getCompaniesCollection() *mongo.Collection {
	databaseNameEnv := os.Getenv(databaseName)
	collectionNameEnv := os.Getenv(collectionName)

	if databaseNameEnv == "" || collectionNameEnv == "" {
		log.Fatal("databaseNameEnv or collectionNameEnv environment variable not set")
	}

	return client.Database(databaseNameEnv).Collection(collectionNameEnv)
}
