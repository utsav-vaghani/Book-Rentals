package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//CommentController struct
type CommentController struct {
	commentRepo *repo.CommentRepository
}

//NewCommentController new book controller
func NewCommentController(db *mongo.Database) *CommentController {
	return &CommentController{
		commentRepo: repo.GetCommentRepository(db),
	}
}

//FetchComments fetch comments of book
func (c *CommentController) FetchComments(ctx *gin.Context) {
	bookID, exists := ctx.Get("book_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "BookID not found"})
		return
	}

	comments, err := c.commentRepo.FetchCommentsByID(bookID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comments": comments, "message": "Orders Fetched Successfully"})
	}
}

//AddComment add comment to book
func (c *CommentController) AddComment(ctx *gin.Context) {
	bookID, exists := ctx.Get("book_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "BookID not found"})
		return
	}

	var comment models.Comment
	_ = ctx.BindJSON(&comment)

	if comment.UserID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request to add comment"})
		return
	}

	_, err := c.commentRepo.AddComment(bookID.(string), comment)

	if err == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add comment to book"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

//RemoveComment from comments of book
func (c *CommentController) RemoveComment(ctx *gin.Context) {
	bookID, exists1 := ctx.Get("book_id")
	commentID, exists2 := ctx.Get("comment_id")

	if !exists1 || !exists2 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "BookID or CommentID not found"})
		return
	}

	err := c.commentRepo.RemoveComment(bookID.(string), commentID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to remove comment!"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment remove successfully"})
	}
}
