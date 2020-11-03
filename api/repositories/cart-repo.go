package repo

import (
	"../../config"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CartRepository struct
type CartRepository struct {
	db *mongo.Collection
}

//GetCartRepository get cart repository
func GetCartRepository(db *mongo.Database) *CartRepository {
	return &CartRepository{
		db: db.Collection(config.Cart),
	}
}

//FetchCart fetch cart by its ID
func (c *CartRepository) FetchCart(userID string) (models.Cart, error) {
	var cart models.Cart
	err := c.db.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&cart)

	return cart, err
}

//AddBook new cart
func (c *CartRepository) AddBook(userID string, book models.Book) (models.Cart, error) {
	var cart = models.Cart{}
	err := c.db.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&cart)

	opts := options.FindOneAndUpdate().SetUpsert(true)

	update := bson.M{
		"total_amount": cart.TotalAmount + float64(book.Price),
		"$push":        bson.M{"cart.$.books": book},
	}

	err = c.db.FindOneAndUpdate(context.TODO(), cart, update, opts).Decode(&cart)

	return cart, err
}

//RemoveBook book from cart
func (c *CartRepository) RemoveBook(userID string, book models.Book) error {
	var cart = models.Cart{}
	err := c.db.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&cart)

	opts := options.FindOneAndUpdate().SetUpsert(true)

	update := bson.M{
		"total_amount": cart.TotalAmount - float64(book.Price),
		"$pull":        bson.M{"cart.$.books": book},
	}

	err = c.db.FindOneAndUpdate(context.TODO(), cart, update, opts).Decode(&cart)

	return err
}
