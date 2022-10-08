package models

import (
	"context"
	"time"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/iamcaleberic/semi-task/utils"
)

// create Mongo client
func CreateDBClient() *mongo.Client {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Connection URI
	uri := utils.CheckEnv("MONGODB_URI")

	// Create a new client and connect to the server
	// TODO: improve retry logic
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	log.Info("Database:Successfully connected and pinged.")

	return client
}

var DBClient *mongo.Client = CreateDBClient()

// returns collection
func GetCollection(collectionName string) *mongo.Collection {

	dbName := utils.CheckEnv("MONGODB_DATABASE")

	collection := DBClient.Database(dbName).Collection(collectionName)
	return collection
}
