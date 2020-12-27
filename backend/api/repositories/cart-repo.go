package repo

import (
	"context"
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
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
func (c *CartRepository) AddBook(books models.Books) (models.Cart, error) {
	var cart models.Cart

	filter := bson.M{
		"user_id": books.UserID,
	}

	update := bson.M{
		"$inc": bson.D{
			{"total_amount", float64(books.Quantity) * books.Price},
		},
		"$push": bson.M{
			"books": books,
		},
	}

	opts := options.FindOneAndUpdate().SetUpsert(true)

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&cart)

	return cart, err
}

//RemoveBook book from cart
func (c *CartRepository) RemoveBook(books models.Books) error {
	var cart = models.Cart{}

	filter := bson.M{
		"user_id": books.UserID,
		"books": bson.M{
			"$elemMatch": bson.M{
				"book_id": books.BookID,
			},
		},
	}

	update := bson.M{
		"$inc": bson.D{
			{"total_amount", -float64(books.Quantity) * books.Price},
			{"books.$.quantity", -books.Quantity},
		},
	}

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cart)

	filter = bson.M{
		"user_id": books.UserID,
	}

	update = bson.M{
		"$pull": bson.M{
			"books": bson.M{
				"book_id": books.BookID,
				"quantity": bson.M{
					"$lte": 0,
				},
			},
		},
	}

	err = c.db.FindOneAndUpdate(context.TODO(), filter, update).Decode(&cart)

	return err
}
