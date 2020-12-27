package repo

import (
	"context"
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
	"github.com/ultra-utsav/Book-Rentals/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
	order.Time = time.Now()
	order.ID = utils.GetObjectID()
	_, err := o.db.InsertOne(context.TODO(), order)

	return err
}
