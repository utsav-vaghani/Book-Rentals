package repo

import (
	"../../config"
	"../../utils"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//CommentRepository struct
type CommentRepository struct {
	db *mongo.Collection
}

//GetCommentRepository get comment repository
func GetCommentRepository(db *mongo.Database) *CommentRepository {
	return &CommentRepository{
		db: db.Collection(config.Comment),
	}
}

//FetchCommentsByID fetch comments of a book
func (c *CommentRepository) FetchCommentsByBookID(bookID string) (models.Comments, error) {
	var comments models.Comments
	err := c.db.FindOne(context.TODO(), bson.M{"book_id": bookID}).Decode(&comments)

	return comments, err
}

//AddComment add comment
func (c *CommentRepository) AddComment(comment models.Comment) error {
	opts := options.FindOneAndUpdate().SetUpsert(true)

	var updatedComments models.Comments

	filter := bson.M{
		"book_id": comment.BookID,
	}

	comment.ID = utils.GetObjectID()

	update := bson.M{
		"$push": bson.M{
			"comments": comment,
		},
	}

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedComments)

	return err
}

//RemoveComment remove comment
func (c *CommentRepository) RemoveComment(bookID, commentID string) error {
	var comments = models.Comments{}

	filter := bson.M{
		"book_id": bookID,
	}

	update := bson.M{
		"$pull": bson.M{
			"comments": bson.M{
				"id": commentID,
			},
		},
	}

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update).Decode(&comments)

	return err
}
