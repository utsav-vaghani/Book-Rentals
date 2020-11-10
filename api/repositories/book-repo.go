package repo

import (
	"../../config"
	"../../utils"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//BookRepository struct
type BookRepository struct {
	db *mongo.Collection
}

//GetBookRepository get new BookRepository
func GetBookRepository(db *mongo.Database) *BookRepository {
	return &BookRepository{
		db: db.Collection(config.Book),
	}
}

//CreateBook to create new book
func (b *BookRepository) CreateBook(book models.Book) (error, bool) {
	var tempBook models.Book
	err := b.db.FindOne(context.TODO(), bson.M{"title": book.Title, "author": book.Author, "owner_id": book.OwnerID}).Decode(&tempBook)

	if err == mongo.ErrNoDocuments {
		book.ID = utils.GetObjectID()
		_, err = b.db.InsertOne(context.TODO(), book)

		if err == nil {
			return nil, true
		}
	}

	return err, false
}

//UpdateBook update book
func (b *BookRepository) UpdateBook(book models.Book) (*mongo.UpdateResult, error) {
	filter := bson.M{
		"id": book.ID,
	}

	update := bson.M{
		"$set": book,
	}

	res, err := b.db.UpdateOne(context.TODO(), filter, update)

	return res, err
}

//FetchBookByID fetch book of given id
func (b *BookRepository) FetchBookByID(id string) (models.Book, error) {
	var book models.Book

	filter := bson.M{
		"id": id,
	}

	err := b.db.FindOne(context.TODO(), filter).Decode(&book)

	return book, err
}

//FetchBooks fetch all books from DB
func (b *BookRepository) FetchBooks() ([]models.Book, error) {
	var books []models.Book

	curr, err := b.db.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}

	err = curr.All(context.TODO(), &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}
