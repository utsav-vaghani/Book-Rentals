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
	commentRepo *repo.CommentRepository
}

//NewBookController new book controller
func NewBookController(db *mongo.Database) *BookController {
	return &BookController{
		bookRepo:    repo.GetBookRepository(db),
		commentRepo: repo.GetCommentRepository(db),
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

//FetchComments fetch comments of book
func (b *BookController) FetchComments(ctx *gin.Context) {
	bookID, exists := ctx.Get("book_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "BookID not found"})
		return
	}

	comments, err := b.commentRepo.FetchCommentsByID(bookID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comments": comments, "message": "Orders Fetched Successfully"})
	}
}

//AddComment add comment to book
func (b *BookController) AddComment(ctx *gin.Context) {
	bookID, exists := ctx.Get("book_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "BookID not found"})
		return
	}

	var comment models.Comment
	ctx.BindJSON(&comment)

	if comment.UserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request to add comment"})
		return
	}

	_, err := b.commentRepo.AddComment(bookID.(string), comment)

	if err == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add comment to book"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

//RemoveComment from comments of book
func (b *BookController) RemoveComment(ctx *gin.Context) {
	bookID, exists1 := ctx.Get("book_id")
	commentID, exists2 := ctx.Get("comment_id")

	if !exists1 || !exists2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "BookID or CommentID not found"})
		return
	}

	err := b.commentRepo.RemoveComment(bookID.(string), commentID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to remove comment!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment remove successfully"})
	}
}
