package updater

import (
	"context"
	"fmt"
	"log"

	//"go.mongodb.org/mongo-driver/bson"
	"javascriptquizgame/parser"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB will connect to the DB
func UpdateDB() bool {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return false
	}

	err = client.Ping(context.TODO(), nil)
	db := client.Database("Tutorials").Collection("Content")

	if err != nil {
		return false
	}
	
	return true	
}