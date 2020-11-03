package main

import (
	"./config"
	"./utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"

	"./api/controllers"
)

func main() {
	router := gin.Default()

	//DB Connection
	db, er := utils.GetConnection()
	if er != nil {
		log.Panic(er)
	}

	//Controllers
	authController := controllers.NewAuthController(db)
	bookController := controllers.NewBookController(db)
	cartController := controllers.NewCartController(db)
	commentController := controllers.NewCommentController(db)
	orderController := controllers.NewOrderController(db)

	//User Routes
	router.POST("/auth/register", authController.RegisterUser)
	router.POST("/auth/login", authController.LoginUser)
	router.GET("/auth/authenticate", authController.AuthenticateUser)

	//Book Routes
	router.POST("/book/create", bookController.CreateBook)
	router.POST("/book/update", bookController.UpdateBook)
	router.POST("/book", bookController.FetchBooks)

	//Cart Routes
	router.POST("/cart/addBook", cartController.AddBookToCart)
	router.POST("/cart/removeBook", cartController.RemoveBookFromCart)
	router.POST("/cart", cartController.FetchCart)

	//Comment Routes
	router.POST("/comment/create", commentController.AddComment)
	router.POST("/comment/delete", commentController.RemoveComment)
	router.POST("/comment", commentController.FetchComments)

	//Order Routes
	router.POST("/order/create", orderController.NewOrder)
	router.POST("/order", orderController.FetchOrders)

	//No Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not Found!")
	})

	// server listening
	router.Run(":" + config.PORT)
}
