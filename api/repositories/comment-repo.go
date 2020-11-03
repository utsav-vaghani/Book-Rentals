package repo

import (
	"../../config"
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
func (c *CommentRepository) FetchCommentsByID(bookID string) (models.Comments, error) {
	var comments models.Comments
	err := c.db.FindOne(context.TODO(), bson.M{"book_id": bookID}).Decode(&comments)

	return comments, err
}

//AddComment add comment
func (c *CommentRepository) AddComment(bookID string, comment models.Comment) (models.Comments, error) {
	opts := options.FindOneAndUpdate().SetUpsert(true)

	var updatedComments models.Comments
	filter := bson.M{"book_id": bookID}
	update := bson.M{"$push": bson.M{"comments.$.comments": comment}}

	err := c.db.FindOneAndUpdate(context.TODO(), filter, update, opts).Decode(&updatedComments)

	return updatedComments, err
}

//RemoveComment remove comment
func (c *CommentRepository) RemoveComment(bookID, commentID string) error {
	var comments = models.Comments{}
	err := c.db.FindOne(context.TODO(), bson.M{"book_id": bookID}).Decode(&comments)

	if err != nil {
		return err
	}

	opts := options.FindOneAndUpdate().SetUpsert(true)

	update := bson.M{
		"$pull": bson.M{"comments.$.comments._id": commentID},
	}

	err = c.db.FindOneAndUpdate(context.TODO(), comments, update, opts).Decode(&comments)

	return err
}
