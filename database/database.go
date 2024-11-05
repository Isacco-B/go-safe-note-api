package database

import (
    "context"
    "fmt"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func ConnectDatabase() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	fmt.Println("Connected to MongoDB!")
	Client = client
	return nil
}

func GetCollection(databaseName, collectionName string) (*mongo.Collection, error) {
	if Client == nil {
		return nil, fmt.Errorf("database not connected")
	}
	return Client.Database(databaseName).Collection(collectionName), nil
}
