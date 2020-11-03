package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//CartController struct
type CartController struct {
	cartRepo *repo.CartRepository
}

//NewCartController new user controller
func NewCartController(db *mongo.Database) *CartController {
	return &CartController{
		cartRepo: repo.GetCartRepository(db),
	}
}

//FetchCart fetch cart of a user
func (u *CartController) FetchCart(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "UserID not found"})
	}

	cart, err := u.cartRepo.FetchCartByID(userID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch cart"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"cart": cart, "message": "Cart Fetched Successfully"})
	}
}

//AddBookToCart add book to the cart
func (u *CartController) AddBookToCart(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
	}

	var book models.Book
	_ = ctx.BindJSON(&book)

	cart, err := u.cartRepo.AddBook(userID.(string), book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add book to the cart"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Book added Successfully to the cart!", "Cart": cart})
	}
}
