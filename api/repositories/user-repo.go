package repo

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	db *mongo.Database
}

func GetUserRepository(db *mongo.Database)*UserRepository{
	return &UserRepository{db: db}
}

func (u *UserRepository) RegisterUser()  {

}