package main

import (
	"./config"
	"./utils"
	"github.com/gin-gonic/gin"
	"log"

	"./api/controllers"
)

func main() {
	router := gin.Default()
	db, er := utils.GetConnection()
	if er != nil {
		log.Panic(er)
	}

	//Initialize context of controllers
	userController := controllers.GetUserController(db)

	//Apply routes
	router.POST("/user/register",userController.RegisterUser)
	router.NoRoute( func(context *gin.Context) {
		context.Writer.WriteString("404 Not Found!")
	})

	// server listening
	router.Run(":" + config.PORT)
}
