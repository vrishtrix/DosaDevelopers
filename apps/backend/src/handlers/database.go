package handlers

import (
	"context"
	"fmt"
	"log"
	"dosadevelopers.devsoc/backend/src/config"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


var usercollection *mongo.Collection
var client *mongo.Client

func ConnectToDB() {
	connectionurl := config.MONGO_URL
	clientOptions := options.Client().ApplyURI(connectionurl)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// Create a MongoDB client
	client, ERR := mongo.Connect(ctx, clientOptions)
	if ERR != nil {
		log.Fatal(ERR)
	}

	// Set a timeout for the connection
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// Connect to MongoDB Atlas
	if ERR != nil {
		log.Fatal(ERR)
	}

	// Ping the MongoDB server to check the connection
	ERR = client.Ping(ctx, nil)
	if ERR != nil {
		log.Fatal(ERR)
	}

	fmt.Println("Connected to MongoDB Atlas!")

	usercollection = client.Database("unmute").Collection("user")

}

func DisconnectFromDB() {
	// Disconnect from MongoDB Atlas
	err := client.Disconnect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB Atlas!")
}
