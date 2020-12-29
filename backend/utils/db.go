package utils

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

//GetConnection get new connection of MongoDB
func GetConnection() (*mongo.Database, error) {
	//select client options
	clientOptions := options.Client().ApplyURI(config.MongoURI)

	//connect to MongoDB
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Println("Error in connecting to MongoDB" + err.Error())
		return nil, err
	}

	log.Println("Connected to MongoDB : " + config.MongoURI)

	return mongoClient.Database(config.Database), nil
}

//InitRedis init connection to redis
func InitRedis() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{Addr: config.RedisDsn})

	_, err := client.Ping().Result()

	return client, err
}
