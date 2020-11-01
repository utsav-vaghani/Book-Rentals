package controllers

import (
	"../dtos"
	"../models"
	"../repositories"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//AuthController controller for user apis
type AuthController struct {
	userRepo *repo.AuthRepository
}

//NewAuthController new UserController
func NewAuthController(db *mongo.Database) *AuthController {
	return &AuthController{userRepo: repo.GetUserRepository(db)}
}

//RegisterUser New User
func (u *AuthController) RegisterUser(ctx *gin.Context) {
	var user models.User
	ctx.BindJSON(&user)

	er, registered := u.userRepo.Register(user)

	if registered {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully Registered"})
	} else if er != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Error in register user" + er.Error()})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Email Id Already Exist!"})
	}
}

//LoginUser Login user
func (u *AuthController) LoginUser(ctx *gin.Context) {
	var userDto dtos.LoginDto
	ctx.BindJSON(&userDto)

	user := u.userRepo.Login(userDto)
	//jwt token remaining generation remaining
	if user != nil {
		ctx.JSON(http.StatusOK, gin.H{"message": "Login Successfully"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials!"})
	}
}

//AuthenticateUser authenticate user
func (u *AuthController) AuthenticateUser(ctx *gin.Context) {
	//authenticate jwt token
}
