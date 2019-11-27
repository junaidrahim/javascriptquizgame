package database

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"javascriptquizgame/parser"
	"javascriptquizgame/logger"
)


// UpdateDB will connect to the DB
func UpdateDB() bool {
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.WriteLog("database.go : Database Connection Error")
		return false
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		logger.WriteLog("database.go : Database Ping Error")
		return false
	}

	questionsDB := client.Database("javascriptquizgame").Collection("questions")
	questions := parser.GetQuestions()

	_, err = questionsDB.DeleteMany(context.TODO(), bson.M{})

	if err != nil {
		logger.WriteLog("database.go : Bulk Deletion Error")
	}

	for _, q := range questions {
		_ ,err := questionsDB.InsertOne(context.TODO(), q)

		if err != nil {
			logger.WriteLog("database.go : Insertion Error")
		}
	}

	return true
}

// GetQuestions will query the database and ask for the questions
func GetQuestions() []parser.Question {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		logger.WriteLog("database.go : Database Connection Error")
		return nil
	}

	err = client.Ping(context.TODO(), nil);

	if err != nil {
		logger.WriteLog("database.go : Database Ping Error")
		return nil
	}

	questionsDB := client.Database("javascriptquizgame").Collection("questions")
	questionsCursor, err := questionsDB.Find(context.TODO(), bson.D{})

	if err != nil {
		logger.WriteLog("database.go : Question Find Error")
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
		logger.WriteLog("database.go : Database Disconnection Error")
	}

	return q
}