package utils

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetObjectID() string {
	return primitive.NewObjectID().Hex()
}
