package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//BookController struct
type BookController struct {
	bookRepo    *repo.BookRepository
}

//NewBookController new book controller
func NewBookController(db *mongo.Database) *BookController {
	return &BookController{
		bookRepo:    repo.GetBookRepository(db),
	}
}

//CreateBook create new book
func (b *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	ctx.BindJSON(&book)

	err, created := b.bookRepo.CreateBook(book)

	if created {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully Book Created"})
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in creating book" + err.Error()})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Book Already Exist!"})
	}
}

//FetchBooks fetch all books
func (b *BookController) FetchBooks(ctx *gin.Context) {
	books, err := b.bookRepo.FetchBooks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch books"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"books": books, "message": "Books Fetched Successfully"})
	}
}

