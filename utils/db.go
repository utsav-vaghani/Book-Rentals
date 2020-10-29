package utils

import (
	"../config"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//GetConnection get new connection of MongoDB
func GetConnection() (*mongo.Database,error) {
	//select client options
	clientOptions := options.Client().ApplyURI(config.MongoURI)

	//connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(),clientOptions)

	if err != nil {
		log.Println("Error in connecting to MongoDB" + err.Error())
		return nil,err
	}

	log.Println("Connected to MongoDB : "+ config.MongoURI)

	return mongoClient.Database(config.Database),nil
}