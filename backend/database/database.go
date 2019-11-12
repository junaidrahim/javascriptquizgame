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
// TODO : Implement Error handling

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
		}
	}

	return true
}

// GetQuestions will query the database and ask for the questions
func GetQuestions() []parser.Question {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		// TODO: Log some errors
		fmt.Println("Connection Error")
		return nil
	}

	err = client.Ping(context.TODO(), nil);

	if err != nil {
		// TODO: Log some errors
		fmt.Println("Ping error")
		return nil
	}

	questionsDB := client.Database("javascriptquizgame").Collection("questions")
	questionsCursor, err := questionsDB.Find(context.TODO(), bson.D{})

	if err != nil {
		// TODO : log some errors
		fmt.Println("Find error")
		return nil
	}

	var q []parser.Question

	for questionsCursor.Next(context.TODO()) {
		var question parser.Question
		err = questionsCursor.Decode(&question)

		q = append(q, question)
	}

	err = client.Disconnect(context.TODO())

	if err != nil {
		fmt.Println("Disconnection Error")
	}

	return q
}