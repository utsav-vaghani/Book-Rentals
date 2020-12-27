package controllers

import (
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	repo "github.com/ultra-utsav/Book-Rentals/backend/api/repositories"
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
	mp := make(map[string]interface{})
	_ = ctx.Bind(&mp)

	bookID := mp["book_id"].(string)

	if bookID == "" {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "BookID not found!"})
		return
	}

	comments, err := c.commentRepo.FetchCommentsByBookID(bookID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"comments": comments, "message": "Orders Fetched Successfully"})
	}
}

//AddComment add comment to book
func (c *CommentController) AddComment(ctx *gin.Context) {

	var comment models.Comment
	_ = ctx.BindJSON(&comment)

	if comment.UserID == "" || comment.BookID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "UserID not found!"})
		return
	}

	err := c.commentRepo.AddComment(comment)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add comment to book!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

//RemoveComment from comments of book
func (c *CommentController) RemoveComment(ctx *gin.Context) {
	mp := make(map[string]interface{})
	_ = ctx.Bind(&mp)

	bookID := mp["book_id"].(string)
	commentID := mp["comment_id"].(string)

	if bookID == "" || commentID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "BookID or CommentID not found!"})
		return
	}

	err := c.commentRepo.RemoveComment(bookID, commentID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to remove comment!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Comment removed successfully"})
	}
}
