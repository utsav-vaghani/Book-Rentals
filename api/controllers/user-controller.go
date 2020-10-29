package controllers

import (
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

//UserController controller for user apis
type UserController struct {
	userRepo *repo.UserRepository
}

//GetUserController get UserController
func GetUserController(db *mongo.Database) *UserController {
	return &UserController{userRepo: repo.GetUserRepository(db)}
}

//RegisterUser New User
func (u *UserController) RegisterUser(ctx *gin.Context) {

}
