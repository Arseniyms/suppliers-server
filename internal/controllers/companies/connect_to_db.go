package companies

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	mongoUri       = "MONGO_URI"
	databaseName   = "DATABASE_NAME"
	collectionName = "COLLECTION_NAME"
	COMPANIES_PATH = "/companies/"
)

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
