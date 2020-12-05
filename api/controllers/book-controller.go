package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

//BookController struct
type BookController struct {
	bookRepo *repo.BookRepository
}

//NewBookController new book controller
func NewBookController(db *mongo.Database) *BookController {
	return &BookController{
		bookRepo: repo.GetBookRepository(db),
	}
}

//CreateBook create new book
func (b *BookController) CreateBook(ctx *gin.Context) {
	var book models.Book
	_ = ctx.BindJSON(&book)

	err, created := b.bookRepo.CreateBook(book)

	if created {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully Book Created"})
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create book!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Book Already Exist!"})
	}
}

//UpdateBook update book
func (b *BookController) UpdateBook(ctx *gin.Context) {
	var book models.Book
	_ = ctx.BindJSON(&book)

	log.Println(book)

	res, err := b.bookRepo.UpdateBook(book)

	if res.MatchedCount == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "No matching document found!"})
	} else if res.ModifiedCount == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to update book!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Book updated Successfully"})
	}
}

//FetchBooks fetch all books
func (b *BookController) FetchBooks(ctx *gin.Context) {
	books, err := b.bookRepo.FetchBooks()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch books!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"books": books, "message": "Books Fetched Successfully"})
	}
}
