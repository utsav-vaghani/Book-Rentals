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
	userController := controllers.NewUserController(db)
	bookController := controllers.NewBookController(db)
	authController := controllers.NewAuthController(db)

	//User Routes
	router.POST("/auth/register", authController.RegisterUser)
	router.POST("/auth/login", authController.LoginUser)
	router.GET("/auth/authenticate", authController.AuthenticateUser)

	//Book Routes
	router.POST("/book/create", bookController.CreateBook)
	router.POST("/book/fetch", bookController.FetchBooks)

	//User Routes
	router.GET("/user/cart", userController.FetchCart)
	
	//No Route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "Not Found!")
	})

	// server listening
	router.Run(":" + config.PORT)
}
