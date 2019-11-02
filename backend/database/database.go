package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"javascriptquizgame/parser"
)

// TODO : Setup proper logging using some framework

// UpdateDB will connect to the DB
func UpdateDB() bool {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		return false
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	questionsDB := client.Database("javascriptquizgame").Collection("questions")
	questions := parser.GetQuestions()

	_, err = questionsDB.DeleteMany(context.TODO(), bson.M{})

	for _, q := range questions {
		output ,err := questionsDB.InsertOne(context.TODO(), q)

		if err != nil {
			fmt.Println("Error in inserting questions", err)
		} else {
			fmt.Println(output)
		}
	}

	return true
}

// GetQuestions will query the database and ask for the questions
func GetQuestions(){

}