package utils

import (
	"github.com/twinj/uuid"
)

func GetObjectID() string {
	return uuid.NewV4().String()
}
