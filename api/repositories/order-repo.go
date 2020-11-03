package repo

import (
	"../../config"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
func (o *OrderRepository) FetchOrdersByUserID(userID string) ([]models.Order, error) {
	var orders []models.Order

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

//NewOrder place new order
func (o *OrderRepository) NewOrder(order models.Order) error {
	var findOrder models.Order
	err := o.db.FindOne(context.TODO(), order).Decode(&findOrder)

	if err == mongo.ErrNoDocuments {
		_, err = o.db.InsertOne(context.TODO(), order)
	}

	return err
}

//CheckoutOrder checkout order
func (o *OrderRepository) CheckoutOrder(orderID string) error {
	_, err := o.db.DeleteOne(context.TODO(), bson.M{"_id": primitive.ObjectIDFromHex(orderID)})
	return err
}
