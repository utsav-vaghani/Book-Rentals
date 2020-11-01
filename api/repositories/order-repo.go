package repo

import (
	"../../config"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//OrderRepository struct
type OrderRepository struct {
	db *mongo.Collection
}

//GetOrderRepository get order repository
func GetOrderRepository(db *mongo.Database) *OrderRepository {
	return &OrderRepository{
		db: db.Collection(config.Order),
	}
}

//FetchOrdersByID fetch all orders of a user
func (o *OrderRepository) FetchOrdersByUserID(userID string) ([]models.Orders, error) {
	var orders []models.Orders

	curr, err := o.db.Find(context.TODO(), bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}

	err = curr.All(context.TODO(), &orders)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
