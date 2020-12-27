package controllers

import (
	"../dtos"
	"../models"
	"../repositories"
	"fmt"
	"github.com/Book-Rentals/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
	"time"
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
	_ = ctx.BindJSON(&user)

	er, registered := u.userRepo.Register(user)

	if registered {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Successfully Registered"})
	} else if er != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to register user!", "error": er.Error()})
	} else {
		ctx.JSON(http.StatusConflict, gin.H{"message": "Email Id Already Exist!"})
	}
}

//LoginUser Login user
func (u *AuthController) LoginUser(ctx *gin.Context) {
	var loginDto dtos.LoginDto
	_ = ctx.BindJSON(&loginDto)

	if user := u.userRepo.Login(loginDto); user != nil {
		userDto := dtos.MapUserToUserDto(user)
		if token, err := createToken(userDto); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "token": token})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login"})
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials!"})
	}
}

//AuthenticateUser authenticate user
func (u *AuthController) AuthenticateUser(ctx *gin.Context) {
	if userDto, err := extractTokenData(ctx.Request); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"user": userDto})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token!"})
	}
}

//CreateToken to generate new token
func createToken(userDto *dtos.UserDto) (string, error) {
	claims := jwt.MapClaims{}

	claims["authorized"] = true
	claims["id"] = userDto.ID
	claims["name"] = userDto.Name
	claims["email"] = userDto.Email
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return at.SignedString([]byte(config.ACCESS_SECRET))
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	token := strings.Split(bearToken, " ")

	if len(token) == 2 {
		return token[1]
	}

	return ""
}

func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.ACCESS_SECRET), nil
	})

	return token, err
}

func extractTokenData(r *http.Request) (*dtos.UserDto, error) {
	token, err := verifyToken(r)
	fmt.Println(token, err)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id := claims["id"].(string)
		name := claims["name"].(string)
		email := claims["email"].(string)

		if id != "" || name != "" || email != "" {

			return &dtos.UserDto{
				ID:    id,
				Name:  name,
				Email: email,
			}, err
		}
	}

	return nil, err
}
