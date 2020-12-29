package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/ultra-utsav/Book-Rentals/backend/api/dtos"
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	repo "github.com/ultra-utsav/Book-Rentals/backend/api/repositories"
	"github.com/ultra-utsav/Book-Rentals/backend/api/services"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

//AuthController controller for user apis
type AuthController struct {
	userRepo *repo.AuthRepository
	authSvc  *services.AuthService
}

//NewAuthController new UserController
func NewAuthController(db *mongo.Database, client *redis.Client) *AuthController {
	return &AuthController{userRepo: repo.GetUserRepository(db), authSvc: services.GetAuthService(client)}
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
		if token, err := u.authSvc.CreateToken(userDto); err == nil {
			if err = u.authSvc.CreateAuth(userDto.ID, token); err == nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "Login Successfully", "access_token": token.AccessToken, "refresh_token": token.RefreshToken})
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to login"})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials!"})
	}
}

//Logout logout user
func (u *AuthController) LogoutUser(ctx *gin.Context) {
	id, accessUUID := u.authSvc.ExtractTokenMetaData(ctx.Request)
	if len(id) == 0 || len(accessUUID) == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized!"})
		return
	}

	del, err := u.authSvc.DelAuth(accessUUID)
	if err != nil || del == 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized!"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Logout Successfully"})
}

//AuthenticateUser authenticate user
func (u *AuthController) AuthenticateUser(ctx *gin.Context) {
	if userDto, err := u.authSvc.ExtractTokenData(ctx.Request); err == nil {
		ctx.JSON(http.StatusOK, gin.H{"user": userDto})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Token!"})
	}
}

func (u *AuthController) RefreshToken(ctx *gin.Context) {
	tokenMap := make(map[string]string)

	err := ctx.ShouldBindJSON(&tokenMap)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to refresh token!"})
		return
	}

	refreshToken := tokenMap["refresh_token"]

	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing Method")
		}
		return []byte(config.RefreshSecret), nil
	})

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh Token Expired!"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		refreshUUID, ok := claims["refresh_uuid"].(string)
		if !ok {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": "Unable to refresh token!"})
			return
		}

		del, err := u.authSvc.DelAuth(refreshUUID)

		if del == 0 || err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to delete refresh auth!"})
			return
		}

		var userDto *dtos.UserDto
		id := claims["id"].(string)
		name := claims["name"].(string)
		email := claims["email"].(string)

		if id != "" || name != "" || email != "" {
			userDto.ID = id
			userDto.Email = email
			userDto.Name = name
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create new tokens!"})
			return
		}

		ts, err := u.authSvc.CreateToken(userDto)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create new tokens!"})
			return
		}

		err = u.authSvc.CreateAuth(userDto.ID, ts)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Unable to create new tokens!"})
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{"message": "New tokens generated!", "access_token": ts.AccessToken, "refresh_token": ts.RefreshToken})
		return
	}

	ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Refresh Token Expired!"})

}
