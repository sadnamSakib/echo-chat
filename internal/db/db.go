package db

import (
	"context"
	"log"
	"time"

	"github.com/sadnamSakib/echo-chat/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDatabase *mongo.Database

// Connect initializes the MongoDB connection
func Connect() {
	dbConfig := config.AppConfig.Database.MongoDB

	// Create a new client and connect to the server
	clientOptions := options.Client().ApplyURI(dbConfig.URI)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %s", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %s", err)
	}

	// Ping the MongoDB server to check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %s", err)
	}

	log.Println("Connected to MongoDB!")

	MongoClient = client
	MongoDatabase = client.Database(dbConfig.Database)
}

// Disconnect closes the MongoDB connection
func Disconnect() {
	if err := MongoClient.Disconnect(context.Background()); err != nil {
		log.Fatalf("Failed to disconnect MongoDB client: %s", err)
	}
	log.Println("Disconnected from MongoDB.")
}
