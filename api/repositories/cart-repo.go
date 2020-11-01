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

//FetchCartByID fetch cart by its ID
func (c *CartRepository) FetchCartByID(userID string) (models.Cart, error) {
	var cart models.Cart
	err := c.db.FindOne(context.TODO(), bson.M{"user_id": userID}).Decode(&cart)

	return cart, err
}

//AddCart new cart
func (c *CartRepository) AddCart(userID string, comment models.Cart) (models.Cart, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true)

	var updatedCart models.Comments
	filter := bson.M{"user_id": userID}
	update := bson.M{"$push": bson.M{"comments.$.comments": comment}}

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedComments)

	return updatedComments, err
}