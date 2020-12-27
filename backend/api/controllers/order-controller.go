package controllers

import (
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	repo "github.com/ultra-utsav/Book-Rentals/backend/api/repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//OrderController struct
type OrderController struct {
	orderRepo *repo.OrderRepository
	bookRepo  *repo.BookRepository
	cartRepo  *repo.CartRepository
}

//NewOrderController new user controller
func NewOrderController(db *mongo.Database) *OrderController {
	return &OrderController{
		orderRepo: repo.GetOrderRepository(db),
		bookRepo:  repo.GetBookRepository(db),
		cartRepo:  repo.GetCartRepository(db),
	}
}

//FetchOrders fetch orders of a user
func (o *OrderController) FetchOrdersByUserID(ctx *gin.Context) {
	mp := make(map[string]interface{})
	_ = ctx.Bind(&mp)

	userID := mp["user_id"].(string)

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "UserID not found!"})
		return
	}

	orders, err := o.orderRepo.FetchOrdersByUserID(userID)

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

	var err []error
	for i := 0; i < len(order.Books); i++ {
		er := o.bookRepo.UpdateStock(order.Books[i].BookID, order.Books[i].Quantity)
		if er != nil {
			err = append(err, er)
		}
	}

	er := o.orderRepo.NewOrder(order)
	if er != nil {
		err = append(err, er)
	}

	if len(err) != 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to place order!", "error": err})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"message": "Order placed Successfully"})
	}
}
