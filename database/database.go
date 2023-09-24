package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

// ConectarMongoDB estabelece uma conexão com o MongoDB e retorna um cliente MongoDB.
func ConnectToMongoDB() (*mongo.Client, error) {
	if dbClient != nil {
		return dbClient, nil
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dbClient = client
	return dbClient, nil
}
