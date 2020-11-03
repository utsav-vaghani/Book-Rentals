package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//UserController struct
type UserController struct {
	cartRepo  *repo.CartRepository
	orderRepo *repo.OrderRepository
}

//NewUserController new user controller
func NewUserController(db *mongo.Database) *UserController {
	return &UserController{
		cartRepo:  repo.GetCartRepository(db),
		orderRepo: repo.GetOrderRepository(db),
	}
}

//FetchOrders fetch orders of a user
func (u *UserController) FetchOrders(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
	}

	orders, err := u.orderRepo.FetchOrdersByUserID(userID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"orders": orders, "message": "Orders Fetched Successfully"})
	}
}

//FetchCart fetch cart of a user
func (u *UserController) FetchCart(ctx *gin.Context) {
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
func (u *UserController) AddBookToCart(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
	}

	var book models.Book
	ctx.BindJSON(&book)

	cart, err := u.cartRepo.AddBook(userID.(string), book)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to add book to the cart"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Book added Successfully to the cart!", "Cart": cart})
	}
}
