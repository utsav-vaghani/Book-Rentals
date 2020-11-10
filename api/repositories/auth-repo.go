package repo

import (
	"../../config"
	"../../utils"
	"../dtos"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongo.Collection
}

func GetUserRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db: db.Collection(config.User)}
}

//RegisterUser register new user
func (u *AuthRepository) Register(user models.User) (error, bool) {
	var tempUser models.User

	err := u.db.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&tempUser)

	if err == mongo.ErrNoDocuments {
		hashedPass, er := utils.Encrypt(user.Password)
		if err == nil {
			user.Password = hashedPass
			user.ID = utils.GetObjectID()
			_, err = u.db.InsertOne(context.TODO(), user)

			if err == nil {
				return er, true
			}
		}
	}

	return err, false
}

//LoginUser login user
func (u *AuthRepository) Login(user dtos.LoginDto) *models.User {
	var tempUser models.User
	err := u.db.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&tempUser)

	if err == nil && utils.Decrypt(tempUser.Password, user.Password) {
		return &tempUser
	}

	return nil
}
