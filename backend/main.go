package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/itsjamie/gin-cors"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
	"github.com/ultra-utsav/Book-Rentals/backend/utils"

	"github.com/ultra-utsav/Book-Rentals/backend/api/controllers"
)

func main() {
	router := gin.Default()

	// CORS Middleware
	router.Use(cors.Middleware(cors.Config{
		Origins:         "*",
		Methods:         "GET, PUT, POST, DELETE",
		RequestHeaders:  "Origin, Authorization, Content-Type",
		ExposedHeaders:  "",
		Credentials:     true,
		ValidateHeaders: false,
	}))

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
	router.POST("/api/auth/register", authController.RegisterUser)
	router.POST("/api/auth/login", authController.LoginUser)
	router.GET("/api/auth/authenticate", authController.AuthenticateUser)

	//Book Routes
	router.POST("/api/book/create", bookController.CreateBook)
	router.POST("/api/book/update", bookController.UpdateBook)
	router.POST("/api/book", bookController.FetchBooks)

	//Cart Routes
	router.POST("/api/cart/addBook", cartController.AddBookToCart)
	router.POST("/api/cart/removeBook", cartController.RemoveBookFromCart)
	router.POST("/api/cart/", cartController.FetchCart)

	//Comment Routes
	router.POST("/api/comment/create", commentController.AddComment)
	router.POST("/api/comment/delete", commentController.RemoveComment)
	router.POST("/api/comment/", commentController.FetchComments)

	//Order Routes
	router.POST("/api/order/create", orderController.NewOrder)
	router.POST("/api/order", orderController.FetchOrdersByUserID)

	//No Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not Found")
	})

	// server listening
	router.Run(":" + config.PORT)
}
