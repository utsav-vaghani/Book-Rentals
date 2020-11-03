package repo

import (
	"../../config"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		_, err = b.db.InsertOne(context.TODO(), book)

		if err == nil {
			return nil, true
		}
	}

	return err, false
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

//UpdateBook update book
func (b *BookRepository) UpdateBook(book models.Book) error {
	filter := bson.M{
		"_id": book.ID,
	}

	opts := options.FindOneAndUpdate().SetUpsert(true)

	err := b.db.FindOneAndUpdate(context.TODO(), filter, book, opts).Decode(&book)

	return err
}
