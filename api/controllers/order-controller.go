package controllers

import (
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
		ctx.JSON(http.StatusNotFound, gin.H{"message": "ID not found"})
	}

	orders, err := o.orderRepo.FetchOrdersByUserID(userID.(string))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to fetch orders"})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"orders": orders, "message": "Orders Fetched Successfully"})
	}
}
