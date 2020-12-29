package services

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis"
	"github.com/twinj/uuid"
	"github.com/ultra-utsav/Book-Rentals/backend/api/dtos"
	"github.com/ultra-utsav/Book-Rentals/backend/api/models"
	"github.com/ultra-utsav/Book-Rentals/backend/config"
	"net/http"
	"strings"
	"time"
)

type AuthService struct {
	client *redis.Client
}

//GetAuthService authService
func GetAuthService(client *redis.Client) *AuthService {
	return &AuthService{client: client}
}

//CreateToken to generate new token
func (a *AuthService) CreateToken(userDto *dtos.UserDto) (*models.Token, error) {
	token := &models.Token{}

	token.AtExp = time.Now().Add(time.Minute * 10).Unix()
	token.AccessUUID = uuid.NewV4().String()

	token.RtExp = time.Now().Add(time.Hour * 24 * 7).Unix()
	token.RefreshUUID = uuid.NewV4().String()

	//Generate Access Token
	atclaims := jwt.MapClaims{}

	atclaims["authorized"] = true
	atclaims["access_uuid"] = token.AccessUUID
	atclaims["id"] = userDto.ID
	atclaims["name"] = userDto.Name
	atclaims["email"] = userDto.Email
	atclaims["exp"] = token.AtExp

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atclaims)

	if accessToken, err := at.SignedString([]byte(config.AccessSecret)); err != nil {
		token.AccessToken = accessToken
	} else {
		return nil, err
	}

	//Generate Refresh Token

	rtClaims := jwt.MapClaims{}

	rtClaims["authorized"] = true
	rtClaims["refresh_uuid"] = token.RefreshUUID
	rtClaims["id"] = userDto.ID
	rtClaims["name"] = userDto.Name
	rtClaims["email"] = userDto.Email
	rtClaims["exp"] = token.RtExp

	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)

	if refreshToken, err := rt.SignedString([]byte(config.RefreshSecret)); err != nil {
		token.RefreshToken = refreshToken
	} else {
		return nil, err
	}

	return token, nil
}

//ExtractTokenData from token
func (a *AuthService) ExtractTokenData(r *http.Request) (*dtos.UserDto, error) {
	token, err := verifyToken(r)
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

//ExtractTokenData from token
func (a *AuthService) ExtractTokenMetaData(r *http.Request) (string, string) {
	token, err := verifyToken(r)
	if err != nil {
		return "", ""
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		id := claims["id"].(string)
		accessUUID := claims["access_uuid"].(string)

		if id != "" || accessUUID != "" {

			return id, accessUUID
		}
	}

	return "", ""
}

//CreateAuth to save token in redis
func (a *AuthService) CreateAuth(userID string, token *models.Token) error {
	atExp := time.Unix(token.AtExp, 0)
	rtExp := time.Unix(token.RtExp, 0)
	now := time.Now()

	if err := a.client.Set(token.AccessUUID, userID, atExp.Sub(now)).Err(); err != nil {
		return err
	}

	if err := a.client.Set(token.RefreshUUID, userID, rtExp.Sub(now)).Err(); err != nil {
		return err
	}

	return nil
}

//FetchAuth fetch from redis
func (a *AuthService) FetchAuth(accessUUID string) (string, error) {
	return a.client.Get(accessUUID).Result()
}

//DelAuth to delete auth
func (a *AuthService) DelAuth(accessUUID string) (int64, error) {
	auth, err := a.client.Del(accessUUID).Result()
	return auth, err
}

//extractToken extract token from header
func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	token := strings.Split(bearToken, " ")

	if len(token) == 2 {
		return token[1]
	}

	return ""
}

//verifyToken
func verifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := extractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method!")
		}
		return []byte(config.AccessSecret), nil
	})

	return token, err
}
