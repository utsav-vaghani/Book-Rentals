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
	//book, err := u.bookRepo.FetchBookByID(books.BookID)
	//
	//if err == mongo.ErrNoDocuments {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "Book does not exist!"})
	//	return
	//} else if err != nil {
	//	ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add book to the cart!!"})
	//	return
	//}
	//
	//if book.Stock < books.Quantity {
	//	ctx.JSON(http.StatusBadRequest, gin.H{"message": "stock is not available"})
	//	return
	//}
	return err
}

//CheckoutOrder checkout order
func (o *OrderRepository) CheckoutOrder(orderID string) error {
	_id, _ := primitive.ObjectIDFromHex(orderID)
	_, err := o.db.DeleteOne(context.TODO(), bson.M{"_id": _id})
	return err
}
