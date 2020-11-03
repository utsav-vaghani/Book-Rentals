package controllers

import (
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//OrderController struct
type OrderController struct {
	orderRepo *repo.OrderRepository
}

//NewOrderController new user controller
func NewOrderController(db *mongo.Database) *OrderController {
	return &OrderController{
		orderRepo: repo.GetOrderRepository(db),
	}
}

//FetchOrders fetch orders of a user
func (o *OrderController) FetchOrders(ctx *gin.Context) {
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "UserID not found!"})
		return
	}

	orders, err := o.orderRepo.FetchOrdersByUserID(userID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"orders": orders, "message": "Orders Fetched Successfully"})
	}
}

//NewOrder add new order
func (o *OrderController) NewOrder(ctx *gin.Context) {
	var order models.Order
	_ = ctx.BindJSON(&order)

	if order.UserID == "" || len(order.Books) == 0 || order.TotalAmount == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Required field missing to place order!"})
		return
	}

	err := o.orderRepo.NewOrder(order)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to place order!", "error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Order placed Successfully"})
	}
}
