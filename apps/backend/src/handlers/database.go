package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"dosadevelopers.devsoc/backend/src/config"

	_ "github.com/joho/godotenv/autoload"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var usercollection *mongo.Collection
var contractcollection *mongo.Collection
var contractrespcollection *mongo.Collection
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
	contractcollection = client.Database("unmute").Collection("contract")
	contractrespcollection = client.Database("unmute").Collection("contractresponse")

}

func DisconnectFromDB() {
	if client != nil { // Assuming 'client' is your MongoDB client instance variable
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := client.Disconnect(ctx)
		if err != nil {
			fmt.Println("Failed to disconnect from MongoDB:", err)
		} else {
			fmt.Println("Disconnected from MongoDB.")
		}
	} else {
		fmt.Println("MongoDB client is nil, cannot disconnect.")
	}
}
