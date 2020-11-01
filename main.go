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

	//User Routes
	router.POST("/user/register",userController.RegisterUser)
	router.POST("/user/login",userController.LoginUser)
	router.GET("/user/authenticate",userController.AuthenticateUser)

	//Book Routes
	router.POST("/book/create",bookController.CreateBook)
	router.POST("/book/fetch",bookController.FetchBooks)

	//No Route
	router.NoRoute( func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound,"Not Found!")
	})

	// server listening
	router.Run(":" + config.PORT)
}
