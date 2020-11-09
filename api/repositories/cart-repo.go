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
func (c *CartRepository) AddBook(userID string, book models.Books) (models.Cart, error) {
	var cart models.Cart

	filter := bson.M{
		"user_id": userID,
	}

	update := bson.M{
		"$inc": bson.D{
			{"total_amount", book.Price},
		},
		"$push": bson.M{
			"books": book,
		},
	}

	opts := options.FindOneAndUpdate().SetUpsert(true)

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&cart)

	return cart, err
}

//RemoveBook book from cart
func (c *CartRepository) RemoveBook(userID string, book models.Books) error {
	var cart = models.Cart{}

	filter := bson.M{
		"user_id": userID,
		"books":   book,
	}
	update1 := bson.M{
		"$unset": bson.M{
			"books.$": true,
		},
	}

	_ = c.db.FindOneAndUpdate(context.TODO(), filter, update1)

	update := bson.M{
		"$inc": bson.D{
			{"total_amount", -book.Price},
		},
		"$pull": bson.M{
			"books": nil,
		},
	}
	filter = bson.M{
		"user_id": userID,
	}
	err := c.db.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cart)

	return err
}
